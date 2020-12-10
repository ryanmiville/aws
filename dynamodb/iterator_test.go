package dynamodb_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ryanmiville/aws/dynamodb"
	"github.com/ryanmiville/aws/dynamodb/dynamodbfakes"
)

//docs represents pages returned by calls to client.ExecuteStatementWithContext(...)
var docs = [][]map[string]*dyn.AttributeValue{
	{
		{"Field": &dyn.AttributeValue{S: aws.String("hello")}},
		{"Field": &dyn.AttributeValue{S: aws.String("my")}},
		{"Field": &dyn.AttributeValue{S: aws.String("baby")}},
	},
	{
		{"Field": &dyn.AttributeValue{S: aws.String("hello")}},
		{"Field": &dyn.AttributeValue{S: aws.String("my")}},
		{"Field": &dyn.AttributeValue{S: aws.String("darling")}},
	},
}

func TestNext(t *testing.T) {
	t.Run("paginates through all docs", func(t *testing.T) {
		ctx := context.Background()
		client := testClient(docs...)
		iter := dynamodb.Scan(client, "table")
		count := 0
		for iter.Next(ctx) {
			count++
		}
		totalDocsCount := 6
		assert.NoError(t, iter.Err())
		assert.Equal(t, totalDocsCount, count)
	})

	t.Run("stays false", func(t *testing.T) {
		ctx := context.Background()
		client := testClient()
		iter := dynamodb.Scan(client, "table")
		for iter.Next(ctx) {
		}
		assert.False(t, iter.Next(ctx))
	})
}

func TestDocument(t *testing.T) {
	type document struct {
		Field string
	}
	expected := []document{
		{"hello"},
		{"my"},
		{"baby"},
		{"hello"},
		{"my"},
		{"darling"},
	}
	ctx := context.Background()
	client := testClient(docs...)
	iter := dynamodb.Scan(client, "table")
	var got []document
	for iter.Next(ctx) {
		var d document
		assert.NoError(t, iter.Document(&d))
		got = append(got, d)
	}
	assert.ElementsMatch(t, expected, got)
}

func TestAPIError(t *testing.T) {
	expected := errors.New("failed api call")
	client := &dynamodbfakes.FakeDynamoDB{}
	client.ExecuteStatementWithContextReturns(nil, expected)
	iter := dynamodb.Scan(client, "table")
	iter.Next(context.Background())
	assert.True(t, errors.Is(iter.Err(), expected))
}

func testClient(pages ...[]map[string]*dyn.AttributeValue) *dynamodbfakes.FakeDynamoDB {
	c := &dynamodbfakes.FakeDynamoDB{}
	next := "next"
	for i, page := range pages {
		out := &dyn.ExecuteStatementOutput{
			Items:     page,
			NextToken: &next,
		}
		c.ExecuteStatementWithContextReturnsOnCall(i, out, nil)
	}
	c.ExecuteStatementWithContextReturns(&dyn.ExecuteStatementOutput{}, nil)
	return c
}
