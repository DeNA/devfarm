package devicefarm

import (
	"encoding/json"
	"fmt"
)

type Platform string

const (
	PlatformIsIOS     Platform = "IOS"
	PlatformIsAndroid Platform = "ANDROID"
)

func (platform Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(platform))
}

func (platform *Platform) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "IOS":
		*platform = PlatformIsIOS
	case "ANDROID":
		*platform = PlatformIsAndroid
	default:
		return fmt.Errorf("unknown platform: %q", s)
	}
	return nil
}
