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
