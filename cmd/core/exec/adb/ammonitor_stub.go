package adb

import (
	"context"
	"io"
)

func StubAmMonitor(errCh <-chan error) ActivityMonitor {
	return func(_ context.Context, _ SerialNumber, _ io.Reader, _ io.Writer, _ io.Writer) error {
		return <-errCh
	}
}

func FakeAmMonitor(err error) ActivityMonitor {
	return func(ctx context.Context, _ SerialNumber, _ io.Reader, _ io.Writer, _ io.Writer) error {
		<-ctx.Done()
		return err
	}
}
