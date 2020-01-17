package awscli

import (
	"bytes"
)

type Version string

type VersionGetter func() (Version, error)

func NewVersionGetter(awsCmd Executor) VersionGetter {
	return func() (Version, error) {
		result, err := awsCmd("--version")
		if err != nil {
			return "", err
		}

		return Version(bytes.TrimSpace(result.Stdout)), nil
	}
}
