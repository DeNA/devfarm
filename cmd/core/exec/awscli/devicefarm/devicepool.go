package devicefarm

import (
	"encoding/json"
	"fmt"
)

type DevicePoolARN string

func (a DevicePoolARN) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a))
}

func (a *DevicePoolARN) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*a = DevicePoolARN(s)
	return nil
}

type DevicePool struct {
	ARN         DevicePoolARN    `json:"arn"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Rules       []DevicePoolRule `json:"rules"`
}

func NewDevicePool(devicePoolARN DevicePoolARN, name string, desc string, rules []DevicePoolRule) DevicePool {
	return DevicePool{
		ARN:         devicePoolARN,
		Name:        name,
		Description: desc,
		Rules:       rules,
	}
}

type DevicePoolRule struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

func NewDeviceARNBasedDevicePoolRule(deviceARN DeviceARN) DevicePoolRule {
	// XXX: In aws-cil documents, the ARN field can take EQUALS as the operator, but it raised "internal error"s.
	//      But it can take IN operator with a single element array, so we should use the way.
	return DevicePoolRule{
		Attribute: "ARN",
		Operator:  "IN",
		Value:     fmt.Sprintf(`[%q]`, string(deviceARN)),
	}
}
