package devicefarm

import (
	"encoding/json"
)

type Availability string

const (
	AvailabilityIsHighlyAvailable       Availability = "HIGHLY_AVAILABLE"
	AvailabilityIsAvailable             Availability = "AVAILABLE"
	AvailabilityIsBusy                  Availability = "BUSY"
	AvailabilityIsTemporaryNotAvailable Availability = "TEMPORARY_NOT_AVAILABLE"
	AvailabilityIsUnknown               Availability = ""
)

func (a Availability) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a))
}

func (a *Availability) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "HIGHLY_AVAILABLE":
		*a = AvailabilityIsHighlyAvailable
	case "AVAILABLE":
		*a = AvailabilityIsAvailable
	case "BUSY":
		*a = AvailabilityIsBusy
	case "TEMPORARY_NOT_AVAILABLE":
		*a = AvailabilityIsTemporaryNotAvailable
	default:
		// XXX: Support Job's device by same struct.
		*a = AvailabilityIsUnknown
	}
	return nil
}
