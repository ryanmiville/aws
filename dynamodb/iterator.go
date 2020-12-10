package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
)

type Iterator struct {
	client     DynamoDB
	input      *dyn.ExecuteStatementInput
	items      []map[string]*dyn.AttributeValue
	curr       map[string]*dyn.AttributeValue
	err        error
	nextCalled bool
	done       bool
}

func Scan(client DynamoDB, table string) *Iterator {
	s := fmt.Sprintf(`SELECT * FROM "%s"`, table)
	return NewIter(client, s)
}

func NewIter(client DynamoDB, statement string) *Iterator {
	input := &dyn.ExecuteStatementInput{Statement: &statement}
	return &Iterator{
		client: client,
		input:  input,
	}
}

func (i *Iterator) Next(ctx context.Context) bool {
	if i.done {
		return false
	}
	if len(i.items) > 0 {
		i.curr = i.items[0]
		i.items = i.items[1:]
		return true
	}
	// we've already called Next and it did not have a next token
	if i.input.NextToken == nil && i.nextCalled {
		i.done = true
		return false
	}
	i.nextCalled = true

	resp, err := i.client.ExecuteStatementWithContext(ctx, i.input)
	if err != nil {
		i.err = err
		i.done = true
		return false
	}
	i.input.NextToken = resp.NextToken
	i.items = resp.Items
	return i.Next(ctx)
}

func (i *Iterator) Document(v interface{}) error {
	return dynamodbattribute.UnmarshalMap(i.curr, v)
}

func (i *Iterator) Err() error {
	return i.err
}
