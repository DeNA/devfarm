package devicefarm

import "encoding/json"

type DevicePoolGetter func(devicePoolARN DevicePoolARN) (DevicePool, error)
type getDevicePoolResponse struct {
	DevicePool DevicePool `json:"devicePool"`
}

func NewDevicePoolGetter(deviceFarmCmd Executor) DevicePoolGetter {
	return func(devicePoolARN DevicePoolARN) (DevicePool, error) {
		result, execErr := deviceFarmCmd(
			"get-device-pool",
			"--arn", string(devicePoolARN),
		)
		if execErr != nil {
			return DevicePool{}, execErr
		}

		var response getDevicePoolResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return DevicePool{}, err
		}

		return response.DevicePool, nil
	}
}
