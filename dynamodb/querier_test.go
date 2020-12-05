package dynamodb_test

import (
	"context"
	"errors"
	"testing"

	ex "github.com/aws/aws-sdk-go/service/dynamodb/expression"

	dyn "github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/ryanmiville/aws/dynamodb"
	"github.com/ryanmiville/aws/dynamodb/dynamodbfakes"
	"github.com/stretchr/testify/assert"
)

func TestNewQuerier(t *testing.T) {
	client := &dynamodbfakes.FakeDynamoDBClient{}
	out := &dyn.QueryOutput{
		Items:            []map[string]*dyn.AttributeValue{docs[0][0], docs[1][0]}, //see scanner_test
		LastEvaluatedKey: docs[1][2],
	}
	client.QueryWithContextReturns(out, nil)
	querier := dynamodb.NewQuerier(client, "table", "Field", dynamodb.DefaultUnmarshal)
	b := ex.NewBuilder().WithKeyCondition(ex.KeyEqual(ex.Key("Field"), ex.Value("hello")))
	querier.Expression(b)
	querier.Next(context.Background())
	_, input, _ := client.QueryWithContextArgsForCall(0)
	assert.Equal(t, "Field", *input.IndexName)
}

func TestQueryFails(t *testing.T) {
	client := &dynamodbfakes.FakeDynamoDBClient{}
	expected := errors.New("query failed")
	client.QueryWithContextReturns(nil, expected)
	querier := dynamodb.NewQuerier(client, "table", "Field", dynamodb.DefaultUnmarshal)
	b := ex.NewBuilder().WithKeyCondition(ex.KeyEqual(ex.Key("Field"), ex.Value("hello")))
	querier.Expression(b)
	assert.False(t, querier.Next(context.Background()))
	assert.True(t, errors.Is(querier.Err(), expected))
}
