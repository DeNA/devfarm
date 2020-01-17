package testutil

import (
	"bytes"
	"io"
)

func NewWriteCloserSpy(base io.WriteCloser) *SpyWriteCloser {
	var inherited io.WriteCloser
	if base != nil {
		inherited = base
	} else {
		inherited = NewWriteCloserStub()
	}

	captured := &bytes.Buffer{}

	return &SpyWriteCloser{
		Inherited: inherited,
		Captured:  captured,
		IsClosed:  false,
	}
}

type SpyWriteCloser struct {
	Inherited io.WriteCloser
	Captured  *bytes.Buffer
	IsClosed  bool
}

func (s *SpyWriteCloser) Write(p []byte) (n int, err error) {
	s.Captured.Write(p)
	return s.Inherited.Write(p)
}

func (s *SpyWriteCloser) Close() error {
	s.IsClosed = true
	return s.Inherited.Close()
}
