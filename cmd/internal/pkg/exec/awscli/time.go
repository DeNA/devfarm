package awscli

import (
	"encoding/json"
	"time"
)

type Timestamp struct {
	Raw time.Time
}

func NewTimestamp(unixSec int64) Timestamp {
	return Timestamp{Raw: time.Unix(unixSec, 0)}
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	var i float64

	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	unixSec := int64(i)
	// FIXME: Get unixNano
	var unixNano int64 = 0

	t.Raw = time.Unix(unixSec, unixNano)
	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Raw.Unix())
}
