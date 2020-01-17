package planfile

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Decode(reader io.Reader) (Planfile, error) {
	var unsafePlanfile UnsafePlanFile
	if err := yaml.NewDecoder(reader).Decode(&unsafePlanfile); err != nil {
		return Planfile{}, err
	}

	planFile, validationErr := Validate(unsafePlanfile)
	if validationErr != nil {
		return Planfile{}, validationErr
	}

	return planFile, nil
}
