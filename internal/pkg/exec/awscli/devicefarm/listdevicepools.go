package devicefarm

import "encoding/json"

type DevicePoolLister func(projectARN ProjectARN) ([]DevicePool, error)

type listDevicePoolsResponse struct {
	DevicePools []DevicePool `json:"devicePools"`
}

func NewDevicePoolLister(deviceFarmCmd Executor) DevicePoolLister {
	return func(projectARN ProjectARN) ([]DevicePool, error) {
		result, execErr := deviceFarmCmd(
			"list-device-pools",
			"--arn", string(projectARN),
		)
		if execErr != nil {
			return nil, execErr
		}

		var response listDevicePoolsResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return nil, err
		}

		return response.DevicePools, nil
	}
}
