package devicefarm

import "encoding/json"

type RunGetter func(runARN RunARN) (Run, error)

type getRunResponse struct {
	Run Run `json:"run"`
}

func NewRunGetter(deviceFarmCmd Executor) RunGetter {
	return func(runARN RunARN) (Run, error) {
		result, execErr := deviceFarmCmd(
			"get-run",
			"--arn", string(runARN),
		)
		if execErr != nil {
			return Run{}, execErr
		}

		var response getRunResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return Run{}, err
		}

		return response.Run, nil
	}
}
