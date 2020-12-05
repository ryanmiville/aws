package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
)

type Iter struct {
	client     DynamoDB
	statement  *string
	items      []map[string]*dyn.AttributeValue
	curr       map[string]*dyn.AttributeValue
	nextToken  *string
	err        error
	nextCalled bool
	done       bool
}

func Scanner(client DynamoDB, table string) *Iter {
	s := fmt.Sprintf(`SELECT * FROM "%s"`, table)
	return NewIter(client, s)
}

func NewIter(client DynamoDB, statement string) *Iter {
	return &Iter{
		client:    client,
		statement: &statement,
	}
}

func (i *Iter) Next(ctx context.Context) bool {
	if i.done {
		return false
	}
	if len(i.items) > 0 {
		i.curr = i.items[0]
		i.items = i.items[1:]
		return true
	}
	// we've already called Next and it did not have a next token
	if i.nextToken == nil && i.nextCalled {
		i.done = true
		return false
	}
	i.nextCalled = true

	in := &dyn.ExecuteStatementInput{
		NextToken: i.nextToken,
		Statement: i.statement,
	}
	resp, err := i.client.ExecuteStatementWithContext(ctx, in)
	if err != nil {
		i.err = err
		i.done = true
		return false
	}
	i.nextToken = resp.NextToken
	i.items = resp.Items
	return i.Next(ctx)
}

func (i *Iter) Document(v interface{}) error {
	return dynamodbattribute.UnmarshalMap(i.curr, v)
}

func (i *Iter) Err() error {
	return i.err
}
