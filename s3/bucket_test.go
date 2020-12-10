package s3_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/ryanmiville/aws/s3"
	"github.com/ryanmiville/aws/s3/s3fakes"
)

//go:generate counterfeiter io.ReadCloser

func TestNewReader(t *testing.T) {
	client := &s3fakes.FakeS3{}
	rc := &s3fakes.FakeReadCloser{}
	resp := &awss3.GetObjectOutput{Body: rc}
	client.GetObjectWithContextReturns(resp, nil)
	b := s3.NewBucket(client, "bucketName")
	ctx := context.Background()

	r, err := b.NewReader(ctx, "path/to/blob")
	assert.NoError(t, err)

	assert.Equal(t, rc, r)
}

func TestNewWriter(t *testing.T) {
	client := &s3fakes.FakeS3{}
	b := s3.NewBucket(client, "bucketName")
	ctx := context.Background()

	w, err := b.NewWriter(ctx, "path/to/blob")
	assert.NoError(t, err)
}
