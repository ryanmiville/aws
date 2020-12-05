package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Creator provides the means to create a document in DynamoDB without needing
// to worry about marshalling or building the input
type Creator struct {
	client DynamoDBClient
	table  string
}

// NewCreator returns a new Creator that can create documents in table
func NewCreator(client DynamoDBClient, table string) *Creator {
	return &Creator{
		client: client,
		table:  table,
	}
}

// Create will create item in the DynamoDB table
func (c *Creator) Create(ctx context.Context, item interface{}) error {
	avmap, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	itemInput := &dynamodb.PutItemInput{
		Item:      avmap,
		TableName: &c.table,
	}
	_, err = c.client.PutItemWithContext(ctx, itemInput)
	if err != nil {
		return fmt.Errorf("failed creating item in DynamoDB: %w", err)
	}
	return nil
}
