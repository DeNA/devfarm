package exec

import "os"

type EnvGetter func(envName string) string

func NewEnvGetter() EnvGetter {
	return func(envName string) string {
		return os.Getenv(envName)
	}
}
