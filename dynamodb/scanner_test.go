package dynamodb_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ryanmiville/aws/dynamodb/dynamodbfakes"

	ex "github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ryanmiville/aws/dynamodb"
)

//docs represents pages returned by calls to client.ScanWithContext(...)
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
	ctx := context.Background()
	scanner := dynamodb.NewScanner(testClient(docs...), "table", dynamodb.DefaultUnmarshal)
	count := 0
	for scanner.Next(ctx) {
		count++
	}
	expected := 0
	for _, page := range docs {
		expected += len(page)
	}
	assert.Equal(t, expected, count)
	assert.NoError(t, scanner.Err())
}

func TestClientErrors(t *testing.T) {
	var expected = errors.New("client failed")
	c := &dynamodbfakes.FakeDynamoDBClient{}
	c.ScanWithContextReturns(nil, expected)
	scanner := dynamodb.NewScanner(c, "table", dynamodb.DefaultUnmarshal)
	scanner.Next(context.Background())
	assert.True(t, errors.Is(scanner.Err(), expected))
}

func TestPagination(t *testing.T) {
	ctx := context.Background()
	client := testClient(docs[0])
	table := "table"
	scanner := dynamodb.NewScanner(client, table, dynamodb.DefaultUnmarshal)
	for scanner.Next(ctx) {
	}

	_, input, _ := client.ScanWithContextArgsForCall(1)
	expected := &dyn.ScanInput{
		TableName:         &table,
		ExclusiveStartKey: map[string]*dyn.AttributeValue{"Field": {S: aws.String("baby")}},
	}

	assert.Equal(t, expected, input)
}

func TestDocument(t *testing.T) {
	ctx := context.Background()
	scanner := dynamodb.NewScanner(testClient(docs...), "table", dynamodb.DefaultUnmarshal)
	var got []map[string]*dyn.AttributeValue
	for scanner.Next(ctx) {
		doc := scanner.Document().(map[string]*dyn.AttributeValue)
		got = append(got, doc)
	}
	assert.NoError(t, scanner.Err())

	var expected []map[string]*dyn.AttributeValue
	for _, page := range docs {
		expected = append(expected, page...)
	}
	assert.ElementsMatch(t, got, expected)
}

func TestUnmarshalFuncErrors(t *testing.T) {
	expected := errors.New("unmarshal failed")
	fn := func(map[string]*dyn.AttributeValue) (interface{}, error) {
		return nil, expected
	}
	scanner := dynamodb.NewScanner(testClient(docs...), "table", fn)
	scanner.Next(context.Background())
	assert.True(t, errors.Is(scanner.Err(), expected))
}

func TestStickyError(t *testing.T) {
	expected := errors.New("unmarshal failed")
	callCount := 0
	fn := func(map[string]*dyn.AttributeValue) (interface{}, error) {
		callCount++
		return nil, expected
	}
	scanner := dynamodb.NewScanner(testClient(docs...), "table", fn)
	scanner.Next(context.Background())
	assert.Equal(t, 1, callCount)
	scanner.Next(context.Background())
	assert.Equal(t, 1, callCount)
}

func TestExpression(t *testing.T) {
	ctx := context.Background()

	cond := ex.Name("Field").Equal(ex.Value("hello"))
	b := ex.NewBuilder().WithFilter(cond)

	table := "table"
	client := testClient()
	scanner := dynamodb.NewScanner(client, table, dynamodb.DefaultUnmarshal)
	scanner.Expression(b)
	scanner.Next(ctx)

	expected := &dyn.ScanInput{
		ExpressionAttributeNames: map[string]*string{"#0": aws.String("Field")},
		ExpressionAttributeValues: map[string]*dyn.AttributeValue{
			":0": {S: aws.String("hello")},
		},
		FilterExpression: aws.String("#0 = :0"),
		TableName:        &table,
	}

	_, actual, _ := client.ScanWithContextArgsForCall(0)
	assert.Equal(t, expected, actual)
}

func TestBadExpression(t *testing.T) {
	scanner := dynamodb.NewScanner(testClient(), "table", dynamodb.DefaultUnmarshal)
	b := ex.NewBuilder()
	scanner.Expression(b)
	assert.Error(t, scanner.Err())
}

func testClient(pages ...[]map[string]*dyn.AttributeValue) *dynamodbfakes.FakeDynamoDBClient {
	c := &dynamodbfakes.FakeDynamoDBClient{}
	for i, page := range pages {
		out := &dyn.ScanOutput{
			Items:            page,
			LastEvaluatedKey: page[len(page)-1],
		}
		c.ScanWithContextReturnsOnCall(i, out, nil)
	}
	c.ScanWithContextReturns(&dyn.ScanOutput{
		Items:            nil,
		LastEvaluatedKey: nil,
	}, nil)
	return c
}
