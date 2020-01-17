package awsdevicefarm

import (
	"encoding/base64"
	"encoding/json"
)

type TransportableArgs []string

var _ json.Marshaler = TransportableArgs{}
var _ json.Unmarshaler = &TransportableArgs{}

func (args TransportableArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(args))
}

func (args *TransportableArgs) UnmarshalJSON(b []byte) error {
	var s []string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*args = s
	return nil
}

func EncodeAppArgs(appArgs TransportableArgs) (string, error) {
	jsonBytes, jsonErr := appArgs.MarshalJSON()
	if jsonErr != nil {
		return "", jsonErr
	}

	return base64.URLEncoding.EncodeToString(jsonBytes), nil
}

func DecodeAppArgs(s string) (TransportableArgs, error) {
	jsonBytes, decodeErr := base64.URLEncoding.DecodeString(s)
	if decodeErr != nil {
		return nil, decodeErr
	}

	var args TransportableArgs
	if err := json.Unmarshal(jsonBytes, &args); err != nil {
		return nil, err
	}

	return args, nil
}
