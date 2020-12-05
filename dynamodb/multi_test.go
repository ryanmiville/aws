package dynamodb_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/stretchr/testify/assert"

	"github.com/ryanmiville/aws/dynamodb"

	"github.com/ryanmiville/aws/dynamodb/dynamodbfakes"
)

func TestMultiNext(t *testing.T) {
	i1 := testIter("a", "b", "c")
	i2 := testIter("d", "e", "f")
	i3 := testIter("g", "h", "i")
	expectedCount := 9

	multi := dynamodb.MultiIterator(i1, i2, i3)
	ctx := context.Background()
	count := 0

	for multi.Next(ctx) {
		count++
	}

	assert.Equal(t, expectedCount, count)
	assert.NoError(t, multi.Err())
}

func TestMultiIterFails(t *testing.T) {
	i1 := testIter("a", "b", "c")

	expected := errors.New("iter failed")
	i2 := testIter()
	i2.NextReturns(false)
	i2.ErrReturns(expected)

	i3 := testIter("a", "b", "c")
	expectedCount := 3

	multi := dynamodb.MultiIterator(i1, i2, i3)
	ctx := context.Background()
	count := 0

	for multi.Next(ctx) {
		count++
	}

	assert.Equal(t, expectedCount, count)
	assert.True(t, errors.Is(multi.Err(), expected))
}

func TestMultiDocument(t *testing.T) {
	i1 := testIter("a", "b", "c")
	i2 := testIter("d", "e", "f")
	i3 := testIter("g", "h", "i")
	expected := []testDoc{
		{"a"}, {"b"}, {"c"},
		{"d"}, {"e"}, {"f"},
		{"g"}, {"h"}, {"i"},
	}
	multi := dynamodb.MultiIterator(i1, i2, i3)
	ctx := context.Background()

	var got []testDoc
	for multi.Next(ctx) {
		doc := multi.Document().(testDoc)
		got = append(got, doc)
	}

	assert.NoError(t, multi.Err())
	assert.ElementsMatch(t, expected, got)
}

func TestMultiExpression(t *testing.T) {
	i1 := testIter("a", "b", "c")
	i2 := testIter("d", "e", "f")
	i3 := testIter("g", "h", "i")

	multi := dynamodb.MultiIterator(i1, i2, i3)
	multi.Expression(expression.NewBuilder())

	assert.Zero(t, i1.ExpressionCallCount())
	assert.Zero(t, i2.ExpressionCallCount())
	assert.Zero(t, i3.ExpressionCallCount())
}

type testDoc struct {
	Field string
}

func testIter(values ...string) *dynamodbfakes.FakeIterator {
	iter := &dynamodbfakes.FakeIterator{}
	for i, v := range values {
		iter.NextReturnsOnCall(i, true)
		iter.DocumentReturnsOnCall(i, testDoc{v})
	}
	iter.NextReturns(false)
	return iter
}
