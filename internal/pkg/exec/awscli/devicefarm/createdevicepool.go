package devicefarm

import "encoding/json"

type DevicePoolCreator func(projectARN ProjectARN, name string, description string, rules []DevicePoolRule) (DevicePool, error)

type createDevicePoolResponse struct {
	DevicePool DevicePool `json:"devicePool"`
}

func NewDevicePoolCreator(deviceFarmCmd Executor) DevicePoolCreator {
	return func(projectARN ProjectARN, name string, description string, rules []DevicePoolRule) (DevicePool, error) {
		rulesJson, rulesErr := json.Marshal(rules)
		if rulesErr != nil {
			return DevicePool{}, rulesErr
		}

		result, execErr := deviceFarmCmd(
			"create-device-pool",
			"--project-arn", string(projectARN),
			"--name", name,
			"--description", description,
			"--rules", string(rulesJson),
		)
		if execErr != nil {
			return DevicePool{}, execErr
		}

		var response createDevicePoolResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return DevicePool{}, err
		}

		return response.DevicePool, nil
	}
}
