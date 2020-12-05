package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// Iterator provides a convenient interface for reading documents from a
// DynamoDB operation. Successive calls to the Next method will step
// through the results of the DynamoDB request, fetching more when
// it reaches the end of a page. The specification of the request
// can be defined by the Expression method. Expression should not be called
// after the first call to Next to ensure predictable behavior.
//
// Iterators must provide a hook for specifying an UnmarshalFunc. Each call
// must unmarshal the next document using the provided UnmarshalFunc.
//
// Iterating stops unrecoverably at the first encountered error
//go:generate counterfeiter . Iterator
type Iterator interface {
	// Next advances the Iterator to the next document, which will then
	// be available through the Document method.
	// Next must use an UnmarshalFunc to allow for safe usage of Document.
	// Next returns false when the operation stops, either by reaching the
	// end of the results or an error.
	// After Next returns false, the Err method will return any error that
	// occurred during iterating.
	Next(ctx context.Context) bool
	// Document returns the most recent document generated by a call
	// to Next as an interface{} specified by an UnmarshalFunc.
	// If the previous call to Next returned true, it is safe to do
	// a type assertion on the returned interface{} to the same
	// underlying type a successful call to the UnmarshalFunc returns.
	Document() interface{}
	// Err returns the first error that was encountered by the Iterator.
	Err() error
	// Expression sets the filters and projections to the request.
	// Expression should not be called after iterating has started.
	Expression(b expression.Builder)
}

type iterator struct {
	get       getFunc
	table     string
	index     *string
	unmarshal UnmarshalFunc
	input     *input
	items     []map[string]*dyn.AttributeValue
	curr      int
	currDoc   interface{}
	last      map[string]*dyn.AttributeValue
	err       error
}

func newIterator(get getFunc, table string, unmarshal UnmarshalFunc, input *input) *iterator {
	return &iterator{
		get:       get,
		table:     table,
		unmarshal: unmarshal,
		input:     input,
		last:      make(map[string]*dyn.AttributeValue),
	}
}
func (i *iterator) Next(ctx context.Context) bool {
	if !i.next(ctx) {
		return false
	}
	avmap := i.document()
	doc, err := i.unmarshal(avmap)
	if err != nil {
		i.err = fmt.Errorf("failed to unmarshal from DynamoDB: %w", err)
		return false
	}
	i.currDoc = doc
	return true
}

func (i *iterator) next(ctx context.Context) bool {
	if i.err != nil {
		return false
	}
	if i.curr < len(i.items) {
		return true
	}
	if i.last == nil {
		return false
	}
	if len(i.last) != 0 {
		i.input.ExclusiveStartKey = i.last
	}
	out, err := i.get(ctx, i.input)
	if err != nil {
		i.err = fmt.Errorf("failed DynamoDB request: %w", err)
		return false
	}
	if len(out.Items) == 0 {
		i.last = out.LastEvaluatedKey
		return i.next(ctx)
	}
	i.items = out.Items
	i.last = out.LastEvaluatedKey
	i.curr = 0
	return true
}

func (i *iterator) Document() interface{} {
	return i.currDoc
}

func (i *iterator) document() map[string]*dyn.AttributeValue {
	n := i.items[i.curr]
	i.curr++
	return n
}

func (i *iterator) Err() error {
	return i.err
}

func (i *iterator) Expression(b expression.Builder) {
	expr, err := b.Build()
	if err != nil {
		i.err = fmt.Errorf("failed building expression: %w", err)
		return
	}
	input := &input{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 &i.table,
		IndexName:                 i.index,
	}
	i.input = input
}

type input struct {
	ExpressionAttributeNames  map[string]*string
	ExpressionAttributeValues map[string]*dynamodb.AttributeValue
	KeyConditionExpression    *string
	FilterExpression          *string
	ProjectionExpression      *string
	TableName                 *string
	IndexName                 *string
	ExclusiveStartKey         map[string]*dyn.AttributeValue
}

type output struct {
	Items            []map[string]*dyn.AttributeValue
	LastEvaluatedKey map[string]*dyn.AttributeValue
}

type getFunc func(ctx context.Context, input *input) (output, error)

func getScan(client DynamoDB) getFunc {
	return func(ctx context.Context, input *input) (output, error) {
		in := &dyn.ScanInput{
			ExpressionAttributeNames:  input.ExpressionAttributeNames,
			ExpressionAttributeValues: input.ExpressionAttributeValues,
			FilterExpression:          input.FilterExpression,
			ProjectionExpression:      input.ProjectionExpression,
			TableName:                 input.TableName,
			IndexName:                 input.IndexName,
			ExclusiveStartKey:         input.ExclusiveStartKey,
		}
		out, err := client.ScanWithContext(ctx, in)
		if err != nil {
			return output{}, err
		}
		return output{
			Items:            out.Items,
			LastEvaluatedKey: out.LastEvaluatedKey,
		}, nil
	}
}

func getQuery(client DynamoDB) getFunc {
	return func(ctx context.Context, input *input) (output, error) {
		in := &dyn.QueryInput{
			ExpressionAttributeNames:  input.ExpressionAttributeNames,
			ExpressionAttributeValues: input.ExpressionAttributeValues,
			KeyConditionExpression:    input.KeyConditionExpression,
			FilterExpression:          input.FilterExpression,
			ProjectionExpression:      input.ProjectionExpression,
			TableName:                 input.TableName,
			IndexName:                 input.IndexName,
			ExclusiveStartKey:         input.ExclusiveStartKey,
		}
		out, err := client.QueryWithContext(ctx, in)
		if err != nil {
			return output{}, err
		}
		return output{
			Items:            out.Items,
			LastEvaluatedKey: out.LastEvaluatedKey,
		}, nil
	}
}
