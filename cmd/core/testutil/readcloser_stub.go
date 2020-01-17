package testutil

import (
	"errors"
	"io"
)

func NewErrorReadCloserStub() io.ReadCloser {
	return &ErrorReadCloserStub{}
}

type ErrorReadCloserStub struct{}

func (*ErrorReadCloserStub) Close() error {
	return errors.New("close: EXPECTED_FAILURE")
}

func (*ErrorReadCloserStub) Read(p []byte) (n int, err error) {
	return 0, errors.New("read: EXPECTED_FAILURE")
}
