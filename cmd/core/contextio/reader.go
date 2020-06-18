package contextio

import (
	"context"
	"io"
)

type Reader struct {
	reader io.Reader
	closer io.Closer
	ctx    context.Context
}

var _ io.ReadCloser = &Reader{}

func NewReader(reader io.Reader, closer io.Closer, ctx context.Context) *Reader {
	contextReader := &Reader{ctx: ctx, closer: closer, reader: reader}

	go func() {
		// NOTE: Prevent goroutines leak.
		doneCh := ctx.Done()
		if doneCh != nil {
			<-doneCh
			_ = contextReader.Close()
		}
	}()
	return contextReader
}

func (r *Reader) Read(p []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.reader.Read(p)
}

func (r *Reader) Close() error {
	return r.closer.Close()
}
