package platforms

import "github.com/dena/devfarm/internal/pkg/strutil"

type AndroidVersion string

func (v AndroidVersion) Less(another AndroidVersion) bool {
	return strutil.LessVersion(string(v), string(another))
}
