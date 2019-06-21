package devicefarm

type UploadDeleter func(uploadARN UploadARN) error

func NewUploadDeleter(deviceFarmCmd Executor) UploadDeleter {
	return func(uploadARN UploadARN) error {
		_, execErr := deviceFarmCmd(
			"delete-upload",
			"--arn", string(uploadARN),
		)
		if execErr != nil {
			return execErr
		}

		return nil
	}
}
