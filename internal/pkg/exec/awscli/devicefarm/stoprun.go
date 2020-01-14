package devicefarm

type RunStopper func(runARN RunARN) error

func NewRunStopper(deviceFarmCmd Executor) RunStopper {
	return func(runARN RunARN) error {
		_, execErr := deviceFarmCmd(
			"stop-run",
			"--arn", string(runARN),
		)

		if execErr != nil {
			return execErr
		}

		return nil
	}
}
