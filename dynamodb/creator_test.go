package dynamodb_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/ryanmiville/aws/dynamodb/dynamodbfakes"

	"github.com/stretchr/testify/assert"

	"github.com/ryanmiville/aws/dynamodb"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		name      string
		clientErr error
	}{
		{name: "client succeeds", clientErr: nil},
		{name: "client fails", clientErr: errors.New("put failed")},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			item := struct{ Field string }{"hello"}
			table := "table"
			client := &dynamodbfakes.FakeDynamoDBClient{}
			client.PutItemWithContextReturns(nil, tc.clientErr)
			creator := dynamodb.NewCreator(client, table)
			err := creator.Create(context.Background(), item)

			expectedInput := &dyn.PutItemInput{
				Item:      map[string]*dyn.AttributeValue{"Field": {S: aws.String("hello")}},
				TableName: &table,
			}
			_, input, _ := client.PutItemWithContextArgsForCall(0)

			assert.Equal(t, expectedInput, input)
			assert.True(t, errors.Is(err, tc.clientErr))
		})
	}
}
