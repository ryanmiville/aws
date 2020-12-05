package dynamodb

import (
	"context"

	ex "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type multiIterator struct {
	iters []Iterator
	err   error
}

// MultiIterator returns an Iterator that's the logical concatenation of
// the provided input iterators. A common use case is when a result set
// cannot be obtained by just one query to a table, but rather requires several queries.
// They're iterated sequentially. Once all iterator's Nexts have returned false, Next
// will return false. If any of the iterators return an error, Err will return that error.
//
// There can be no guarantees around the safety of type assertions called on the result of
// Document. It is the caller's responsibility to verify the unmarshalFuncs used in the
// input iterators, and to type assert appropriately if necessary.
//
// Since it does not make sense to set an expression across iterators, Expression is a no-op
func MultiIterator(iters ...Iterator) Iterator {
	return &multiIterator{iters: iters}
}

func (m *multiIterator) Next(ctx context.Context) bool {
	if m.err != nil || len(m.iters) == 0 {
		return false
	}
	if m.iters[0].Next(ctx) {
		return true
	}
	if m.iters[0].Err() != nil {
		m.err = m.iters[0].Err()
		return false
	}
	m.iters = m.iters[1:]
	return m.Next(ctx)
}

func (m *multiIterator) Document() interface{} {
	return m.iters[0].Document()
}

func (m *multiIterator) Err() error {
	return m.err
}

func (m multiIterator) Expression(b ex.Builder) {
}
