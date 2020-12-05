package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// DynamoDB is the interface that provides API calls to DynamoDB
//go:generate counterfeiter . DynamoDB
type DynamoDB dynamodbiface.DynamoDBAPI

func New(p client.ConfigProvider, cfgs ...*aws.Config) DynamoDB {
	return dyn.New(p, cfgs...)
}

// UnmarshalFunc describes how to unmarshal an item from DynamoDB
type UnmarshalFunc func(map[string]*dyn.AttributeValue) (interface{}, error)

// DefaultUnmarshal is a straight pass-through of the item from DynamoDB
func DefaultUnmarshal(avmap map[string]*dyn.AttributeValue) (interface{}, error) {
	return avmap, nil
}
