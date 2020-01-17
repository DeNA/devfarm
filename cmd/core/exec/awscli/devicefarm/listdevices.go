package devicefarm

import "encoding/json"

type DeviceLister func() ([]Device, error)

type listDevicesResponse struct {
	Devices []Device `json:"devices"`
}

func NewDeviceLister(deviceFarmCmd Executor) DeviceLister {
	return func() ([]Device, error) {
		// NOTE: https://docs.aws.amazon.com/cli/latest/reference/devicefarm/list-devices.html
		result, clientErr := deviceFarmCmd("list-devices")
		if clientErr != nil {
			return nil, clientErr
		}

		var response listDevicesResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return nil, err
		}

		return response.Devices, nil
	}
}
