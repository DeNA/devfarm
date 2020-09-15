package exec

import (
	"github.com/dena/devfarm/cmd/core/testutil"
	"os"
)

func StubStatFunc(result os.FileInfo, err error) StatFunc {
	return func(string) (os.FileInfo, error) {
		return result, err
	}
}

func AnySuccessfulStatFunc() StatFunc {
	return StubStatFunc(nil, nil)
}

func AnyFailedStatFunc() StatFunc {
	return StubStatFunc(nil, testutil.AnyError)
}
