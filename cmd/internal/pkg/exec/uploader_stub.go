package exec

import (
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"io"
	"net/http"
)

func StubUploader(err error) Uploader {
	return func(string, func(*http.Request), io.Reader) error {
		return err
	}
}

func AnyUploader() Uploader {
	return AnyFailedUploader()
}

func AnySuccessfulUploader() Uploader {
	return StubUploader(nil)
}

func AnyFailedUploader() Uploader {
	return StubUploader(testutil.AnyError)
}
