package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBClient is the interface that provides API calls to DynamoDB
//go:generate counterfeiter . DynamoDBClient
type DynamoDBClient interface {
	ScanWithContext(ctx aws.Context, input *dyn.ScanInput, opts ...request.Option) (*dyn.ScanOutput, error)
	PutItemWithContext(ctx aws.Context, input *dyn.PutItemInput, opts ...request.Option) (*dyn.PutItemOutput, error)
	QueryWithContext(ctx aws.Context, input *dyn.QueryInput, opts ...request.Option) (*dyn.QueryOutput, error)
}

// UnmarshalFunc describes how to unmarshal an item from DynamoDB
type UnmarshalFunc func(map[string]*dyn.AttributeValue) (interface{}, error)

// DefaultUnmarshal is a straight pass-through of the item from DynamoDB
func DefaultUnmarshal(avmap map[string]*dyn.AttributeValue) (interface{}, error) {
	return avmap, nil
}
