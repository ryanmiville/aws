package s3_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	awss3 "github.com/aws/aws-sdk-go/service/s3"

	"github.com/stretchr/testify/assert"

	"github.com/ryanmiville/aws/s3"
	"github.com/ryanmiville/aws/s3/s3fakes"
)

func TestNewReader(t *testing.T) {
	t.Run("gets a reader from S3", func(t *testing.T) {
		client := &s3fakes.FakeClient{}
		rc := ioutil.NopCloser(&bytes.Buffer{})
		resp := &awss3.GetObjectOutput{Body: rc}
		client.GetObjectWithContextReturns(resp, nil)
		b := s3.NewBucket(client, "bucketName")
		ctx := context.Background()

		r, err := b.NewReader(ctx, "path/to/blob")
		assert.NoError(t, err)

		assert.Equal(t, rc, r)
	})

	t.Run("client fails", func(t *testing.T) {
		client := &s3fakes.FakeClient{}
		expected := errors.New("get object failed")
		client.GetObjectWithContextReturns(nil, expected)
		b := s3.NewBucket(client, "bucketName")
		ctx := context.Background()

		r, err := b.NewReader(ctx, "path/to/blob")
		assert.True(t, errors.Is(err, expected))
		assert.Nil(t, r)
	})
}

func TestNewWriter(t *testing.T) {
	t.Run("writes to s3", func(t *testing.T) {
		client := &s3fakes.FakeClient{}
		var read []byte
		var readErr error
		client.UploadWithContextStub = func(_ context.Context, in *s3manager.UploadInput, opts ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
			read, readErr = ioutil.ReadAll(in.Body)
			return nil, nil
		}
		b := s3.NewBucket(client, "bucketName")
		ctx := context.Background()

		w := b.NewWriter(ctx, "path/to/blob")

		n, err := w.Write([]byte("hello "))
		assert.Equal(t, 6, n)
		assert.NoError(t, err)

		n, err = w.Write([]byte("world"))
		assert.Equal(t, 5, n)
		assert.NoError(t, err)

		err = w.Close()
		assert.NoError(t, err)

		assert.Equal(t, []byte("hello world"), read)
		assert.NoError(t, readErr)
	})

	t.Run("client fails", func(t *testing.T) {
		client := &s3fakes.FakeClient{}
		expected := errors.New("upload failed")
		client.UploadWithContextStub = func(_ context.Context, in *s3manager.UploadInput, opts ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
			ioutil.ReadAll(in.Body)
			return nil, expected
		}
		b := s3.NewBucket(client, "bucketName")
		ctx := context.Background()

		w := b.NewWriter(ctx, "path/to/blob")

		n, err := w.Write([]byte("hello "))
		assert.Equal(t, 6, n)
		assert.NoError(t, err)

		n, err = w.Write([]byte("world"))
		assert.Equal(t, 5, n)
		assert.NoError(t, err)

		err = w.Close()
		assert.True(t, errors.Is(err, expected))
	})

	t.Run("Write after Close", func(t *testing.T) {
		client := &s3fakes.FakeClient{}
		expected := errors.New("sticky error")
		client.UploadWithContextReturns(nil, expected)
		b := s3.NewBucket(client, "bucketName")
		ctx := context.Background()

		w := b.NewWriter(ctx, "path/to/blob")
		_ = w.Close()
		n, err := w.Write([]byte("should not work"))
		assert.True(t, errors.Is(err, expected))
		assert.Zero(t, n)
	})
}
