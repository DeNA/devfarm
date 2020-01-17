package exec

import (
	"github.com/dena/devfarm/cmd/core/testutil"
	"os"
)

func AnyFile() FileStub {
	return FileStub{nextResult: AnyFileStubResult()}
}

func AnySuccessfulFile() FileStub {
	return FileStub{nextResult: AnySuccessfulFileStubResult()}
}

func AnyFileStubResult() FileStubResult {
	return FileStubResult{
		ReadResult:   0,
		ReadError:    testutil.AnyError,
		ReadAtResult: 0,
		ReadAtError:  testutil.AnyError,
		WriteResult:  0,
		WriteError:   testutil.AnyError,
		CloseError:   testutil.AnyError,
		SeekResult:   0,
		SeekError:    testutil.AnyError,
	}
}

func AnySuccessfulFileStubResult() FileStubResult {
	return FileStubResult{
		ReadResult:   0,
		ReadError:    nil,
		ReadAtResult: 0,
		ReadAtError:  nil,
		WriteResult:  0,
		WriteError:   nil,
		CloseError:   nil,
		SeekResult:   0,
		SeekError:    nil,
	}
}

type FileStubResult struct {
	ReadResult   int
	ReadError    error
	ReadAtResult int
	ReadAtError  error
	WriteResult  int
	WriteError   error
	CloseError   error
	SeekResult   int64
	SeekError    error
	StatResult   os.FileInfo
	StatError    error
}

var _ FileLike = &FileStub{}

type FileStub struct {
	nextResult FileStubResult
}

func (f FileStub) Read(p []byte) (int, error) {
	return f.nextResult.ReadResult, f.nextResult.ReadError
}

func (f FileStub) Write(p []byte) (int, error) {
	return f.nextResult.WriteResult, f.nextResult.WriteError
}

func (f FileStub) Close() error {
	return f.nextResult.CloseError
}

func (f FileStub) Seek(offset int64, whence int) (int64, error) {
	return f.nextResult.SeekResult, f.nextResult.SeekError
}

func (f FileStub) Stat() (os.FileInfo, error) {
	return f.nextResult.StatResult, f.nextResult.StatError
}

func (f FileStub) ReadAt(p []byte, off int64) (int, error) {
	return f.nextResult.ReadAtResult, f.nextResult.ReadAtError
}
