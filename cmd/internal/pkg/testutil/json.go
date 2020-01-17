package testutil

import (
	"encoding/json"
	"errors"
	"fmt"
)

func CheckMarshalAndUnmarshalIsEquivalentToOriginal(codable interface{}) error {
	origin := codable

	bytes, marshalErr := json.Marshal(codable)

	if marshalErr != nil {
		return fmt.Errorf("%#v.MarshalJSON() == (nil, %v), but wanted ([]byte, nil)", codable, marshalErr)
	}

	unmarshalErr := json.Unmarshal(bytes, codable)

	if unmarshalErr != nil {
		return fmt.Errorf("json.UnmarshalJSON([]byte(%q), &availability) == %v, but wanted nil", string(bytes), unmarshalErr)
	}

	if codable != origin {
		message := fmt.Sprintf("unmarshal(marshal(%#v)) == %#v, but wanted %#v", origin, codable, origin)
		return errors.New(message)
	}

	return nil
}
