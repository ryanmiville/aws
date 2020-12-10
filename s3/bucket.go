package s3

import (
	"context"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Bucket struct {
	client S3
	name   string
}

func NewBucket(client S3, name string) *Bucket {
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

func (b *Bucket) NewWriter(ctx context.Context, key string) (io.WriteCloser, error) {
	send := func(r io.Reader) error {
		in := &s3manager.UploadInput{
			Bucket: aws.String(b.name),
			Key:    aws.String(key),
			Body:   r,
		}
		_, err := b.client.UploadWithContext(ctx, in)
		return err
	}
	return &writer{
		send: send,
	}, nil
}

type writer struct {
	w     *io.PipeWriter
	send  func(r io.Reader) error
	donec chan struct{}
	err   error
}

// Write appends p to w. User must call Close to close the w after done writing.
func (w *writer) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if w.w == nil {
		pr, pw := io.Pipe()
		w.w = pw
		w.open(pr)
		return 0, nil
	}
	select {
	case <-w.donec:
		return 0, w.err
	default:
	}
	return w.w.Write(p)
}

// pr may be nil if we're Closing and no data was written.
func (w *writer) open(pr *io.PipeReader) {
	go func() {
		defer close(w.donec)
		var r io.Reader = http.NoBody
		if pr != nil {
			r = pr
		}
		err := w.send(r)
		if err != nil {
			w.err = err
			if pr != nil {
				pr.CloseWithError(err)
			}
		}
	}()
}

// Close completes the writer and closes it. Any error occurring during write
// will be returned. If a writer is closed before any Write is called, Close
// will create an empty file at the given key.
func (w *writer) Close() error {
	if w.w == nil {
		// We never got any bytes written. We'll write an http.NoBody.
		w.open(nil)
	} else if err := w.w.Close(); err != nil {
		return err
	}
	<-w.donec
	return w.err
}
