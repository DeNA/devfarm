package platforms

import (
	"fmt"
	"path/filepath"
	"strings"
)

func DetectOSName(path string) (OSName, error) {
	extname := strings.ToLower(filepath.Ext(string(path)))
	switch extname {
	case ".apk":
		return OSIsAndroid, nil
	case ".ipa":
		return OSIsIOS, nil
	default:
		return OSIsUnavailable, fmt.Errorf("cannot detect os by ext: %q", extname)
	}
}
