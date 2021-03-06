package platforms

import (
	"github.com/dena/devfarm/cmd/core/strutil"
)

type IOSVersion string

func (v IOSVersion) Less(another IOSVersion) bool {
	return strutil.LessVersion(string(v), string(another))
}
