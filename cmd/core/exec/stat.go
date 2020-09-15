package exec

import "os"

type StatFunc func(path string) (os.FileInfo, error)

func NewStatFunc() StatFunc {
	return os.Stat
}
