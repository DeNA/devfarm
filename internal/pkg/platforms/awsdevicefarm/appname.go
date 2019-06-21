package awsdevicefarm

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/strutil"
	"path/filepath"
	"strings"
)

type appName string

func validateAppName(appPath ipaOrApkPathOnLocal) (appName, string, error) {
	if appPath == "" {
		return "", "", errors.New("application path must not be empty")
	}

	basename := filepath.Base(string(appPath))
	if err := strutil.IsASCII(basename); err != nil {
		return "", "", fmt.Errorf("basename must contain only ASCII printable chars, but %s", err.Error())
	}

	extname := filepath.Ext(basename)

	return appName(strings.TrimSuffix(basename, extname)), extname, nil
}
