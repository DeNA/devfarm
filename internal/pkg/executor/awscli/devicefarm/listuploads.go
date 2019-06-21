package devicefarm

import "encoding/json"

type UploadLister func(projectARN ProjectARN) ([]Upload, error)

type listUploadsResponse struct {
	Uploads []Upload `json:"uploads"`
}

func NewUploadLister(deviceFarmCmd Executor) UploadLister {
	return func(projectARN ProjectARN) ([]Upload, error) {
		result, execErr := deviceFarmCmd(
			"list-uploads",
			"--arn", string(projectARN),
		)
		if execErr != nil {
			return nil, execErr
		}

		var response listUploadsResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return nil, err
		}

		return response.Uploads, nil
	}
}
