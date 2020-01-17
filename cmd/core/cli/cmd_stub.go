package cli

func AnyCommand() Command {
	return FailureCommand()
}

func SuccessfulCommand() Command {
	return StubCommand(ExitNormal)
}

func FailureCommand() Command {
	return StubCommand(ExitAbnormal)
}

func StubCommand(status ExitStatus) Command {
	return func([]string, ProcessInout) ExitStatus {
		return status
	}
}
