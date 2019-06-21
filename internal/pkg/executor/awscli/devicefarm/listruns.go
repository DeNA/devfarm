package devicefarm

import "encoding/json"

type RunLister func(projectARN ProjectARN) ([]Run, error)

type listRunsResponse struct {
	Runs []Run `json:"runs"`
}

func NewRunLister(deviceFarmCmd Executor) RunLister {
	return func(projectARN ProjectARN) ([]Run, error) {
		result, execErr := deviceFarmCmd(
			"list-runs",
			"--arn", string(projectARN),
		)
		if execErr != nil {
			return nil, execErr
		}

		var response listRunsResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return nil, err
		}

		return response.Runs, nil
	}
}
