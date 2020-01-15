package devicefarm

import "encoding/json"

type UploadGetter func(uploadARN UploadARN) (Upload, error)

type getUploadResponse struct {
	Upload Upload `json:"upload"`
}

func NewUploadGetter(deviceFarmCmd Executor) UploadGetter {
	return func(uploadARN UploadARN) (Upload, error) {
		result, execErr := deviceFarmCmd(
			"get-upload",
			"--arn", string(uploadARN),
		)
		if execErr != nil {
			return Upload{}, execErr
		}

		var response getUploadResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return Upload{}, err
		}

		return response.Upload, nil
	}
}
