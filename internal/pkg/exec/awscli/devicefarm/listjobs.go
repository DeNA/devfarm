package devicefarm

import "encoding/json"

type JobLister func(runARN RunARN) ([]Job, error)

type listJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

func NewJobLister(deviceFarmCmd Executor) JobLister {
	return func(runARN RunARN) ([]Job, error) {
		result, execErr := deviceFarmCmd(
			"list-jobs",
			"--arn", string(runARN),
		)
		if execErr != nil {
			return nil, execErr
		}

		var response listJobsResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return nil, err
		}

		return response.Jobs, nil
	}
}
