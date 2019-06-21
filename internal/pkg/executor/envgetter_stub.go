package executor

func AnyEnvGetter() EnvGetter {
	return StubEnvGetter("ANY_ENV")
}

func StubEnvGetter(envValue string) EnvGetter {
	return func(string) string {
		return envValue
	}
}
