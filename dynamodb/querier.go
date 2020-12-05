package dynamodb

type querier struct {
	*iterator
}

// NewQuerier returns a new Iterator that queries table using idex
// and unmarshals documents using unmarshalFunc. It does not have a
// valid default operation and therefore requires a call to Expression
// before calling Next.
func NewQuerier(client DynamoDB, table string, index string, unmarshalFunc UnmarshalFunc) Iterator {
	input := &input{TableName: &table}
	iter := newIterator(getQuery(client), table, unmarshalFunc, input)
	iter.index = &index
	return &querier{iterator: iter}
}
