package devicefarm

import "encoding/json"

type RunScheduler func(
	projectARN ProjectARN,
	devicePoolARN DevicePoolARN,
	testProp TestProp,
	execConf ExecutionConfiguration,
	appARN UploadARN,
) (Run, error)

type scheduleRunResponse struct {
	Run Run `json:"run"`
}

func NewRunScheduler(deviceFarmCmd Executor) RunScheduler {
	return func(projectARN ProjectARN, devicePoolARN DevicePoolARN, testProp TestProp, execConf ExecutionConfiguration, appARN UploadARN) (Run, error) {
		testPropJSON, testPropErr := json.Marshal(testProp)
		if testPropErr != nil {
			return Run{}, testPropErr
		}

		execConfJSON, execConfErr := json.Marshal(execConf)
		if execConfErr != nil {
			return Run{}, execConfErr
		}

		result, execErr := deviceFarmCmd(
			"schedule-run",
			"--project-arn", string(projectARN),
			"--app-arn", string(appARN),
			"--device-pool-arn", string(devicePoolARN),
			"--test", string(testPropJSON),
			"--execution-configuration", string(execConfJSON),
		)
		if execErr != nil {
			return Run{}, execErr
		}

		var response scheduleRunResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return Run{}, err
		}

		return response.Run, nil
	}
}
