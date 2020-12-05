package dynamodb

type scanner struct {
	*iterator
}

// NewScanner returns a new Iterator that scans table and unmarshals
// documents using unmarshalFunc. It has a default operation of a
// wide open scan of table.
func NewScanner(client DynamoDBClient, table string, unmarshalFunc UnmarshalFunc) Iterator {
	input := &input{TableName: &table}
	iter := newIterator(getScan(client), table, unmarshalFunc, input)
	return &scanner{iterator: iter}
}
