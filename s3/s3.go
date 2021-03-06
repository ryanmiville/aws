package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
)

//go:generate counterfeiter . Client
type Client interface {
	s3iface.S3API
	s3manageriface.UploaderAPI
}

type S3 struct {
	*awss3.S3
	*s3manager.Uploader
}

func New(p client.ConfigProvider, cfgs ...*aws.Config) *S3 {
	s := awss3.New(p, cfgs...)
	u := s3manager.NewUploaderWithClient(s)
	return &S3{
		S3:       s,
		Uploader: u,
	}
}
