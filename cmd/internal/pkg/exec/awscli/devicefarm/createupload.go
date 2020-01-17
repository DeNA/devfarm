package devicefarm

import "encoding/json"

type UploadCreator func(projectARN ProjectARN, uploadFileName UploadFileName, uploadType UploadType) (Upload, error)

type createUploadResponse struct {
	Upload Upload `json:"upload"`
}

func NewUploadCreator(deviceFarmCmd Executor) UploadCreator {
	return func(projectARN ProjectARN, uploadFileName UploadFileName, uploadType UploadType) (Upload, error) {
		// https://docs.aws.amazon.com/cli/latest/reference/devicefarm/create-upload.html
		result, execErr := deviceFarmCmd(
			"create-upload",
			"--project-arn", string(projectARN),
			"--name", string(uploadFileName),
			"--type", string(uploadType),
		)
		if execErr != nil {
			return Upload{}, execErr
		}

		var response createUploadResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return Upload{}, err
		}

		return response.Upload, nil
	}
}
