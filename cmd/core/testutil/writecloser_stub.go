package testutil

import (
	"errors"
	"io"
)

func NewWriteCloserStub() io.WriteCloser {
	return &NullWriteCloser{}
}

type NullWriteCloser struct{}

func (*NullWriteCloser) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (*NullWriteCloser) Close() error {
	return nil
}

func NewErrorClosedWriterStub() io.WriteCloser {
	return &ErrorClosedWriterStub{}
}

type ErrorClosedWriterStub struct{}

func (*ErrorClosedWriterStub) Write(p []byte) (n int, err error) {
	return 0, errors.New("write: EXPECTED_FAILURE")
}

func (*ErrorClosedWriterStub) Close() error {
	return errors.New("close: EXPECTED_FAILURE")
}
