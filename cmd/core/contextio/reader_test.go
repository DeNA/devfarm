package contextio

import (
	"context"
	"io"
	"io/ioutil"
	"testing"
	"time"
)

func TestReaderRead(t *testing.T) {
	internalReader, internalWriter := io.Pipe()
	reader := NewReader(internalReader, internalReader, context.Background())

	go func() {
		time.Sleep(time.Duration(100) * time.Millisecond)
		_, _ = io.WriteString(internalWriter, "Hello")
		_ = internalWriter.CloseWithError(io.EOF)
	}()

	res, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Logf("wrote: %s", res)
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestReaderReadCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	internalReader, _ := io.Pipe()
	reader := NewReader(internalReader, internalReader, ctx)

	// NOTE: Do not use context.WithTimeout because several setup before ReadAll can take some duration.
	go func() {
		time.Sleep(time.Duration(100) * time.Millisecond)
		cancel()
	}()

	_, err := ioutil.ReadAll(reader)
	if err != io.ErrClosedPipe {
		t.Errorf("want io.ErrClosedPipe, got %v", err)
		return
	}
}

func TestReaderCloseAfterCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	internalReader, _ := io.Pipe()
	reader := NewReader(internalReader, internalReader, ctx)

	cancel()

	if err := reader.Close(); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}
