package devicefarm

import (
	"encoding/json"
)

type DeviceARN string

func (a *DeviceARN) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*a = DeviceARN(s)
	return nil
}

func (a DeviceARN) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a))
}

func NewAWSDeviceFarmDevice(arn DeviceARN, manufacturer string, model string, platform Platform, os string, availability Availability) Device {
	return Device{
		ARN:          arn,
		Manufacturer: manufacturer,
		Model:        model,
		Platform:     platform,
		OS:           os,
		Availability: availability,
	}
}

type Device struct {
	ARN          DeviceARN    `json:"arn"`
	Manufacturer string       `json:"manufacturer"`
	Model        string       `json:"model"`
	Platform     Platform     `json:"platform"`
	OS           string       `json:"os"`
	Availability Availability `json:"availability"`
}
