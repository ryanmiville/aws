package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Bucket struct {
	client Client
	name   string
}

func NewBucket(client Client, name string) *Bucket {
	return &Bucket{client: client, name: name}
}

func (b *Bucket) NewReader(ctx context.Context, key string) (io.ReadCloser, error) {
	in := &awss3.GetObjectInput{Bucket: &b.name, Key: &key}
	resp, err := b.client.GetObjectWithContext(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (b *Bucket) NewWriter(ctx context.Context, key string) io.WriteCloser {
	upload := func(r io.Reader) error {
		in := &s3manager.UploadInput{
			Bucket: aws.String(b.name),
			Key:    aws.String(key),
			Body:   r,
		}
		_, err := b.client.UploadWithContext(ctx, in)
		return err
	}
	pr, pw := io.Pipe()
	w := &writer{
		w:     pw,
		donec: make(chan struct{}),
	}
	w.pipe(pr, upload)
	return w
}

type writer struct {
	w     *io.PipeWriter
	donec chan struct{}
	err   error
}

func (w *writer) Write(p []byte) (int, error) {
	select {
	case <-w.donec:
		return 0, w.err
	default:
	}
	return w.w.Write(p)
}

func (w *writer) pipe(pr *io.PipeReader, upload func(r io.Reader) error) {
	go func() {
		defer close(w.donec)
		err := upload(pr)
		if err != nil {
			w.err = err
			_ = pr.CloseWithError(err)
		}
	}()
}

func (w *writer) Close() error {
	_ = w.w.Close()
	<-w.donec
	return w.err
}
