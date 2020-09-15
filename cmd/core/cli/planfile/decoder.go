package planfile

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Decode(planfilePath FilePath, reader io.Reader, validate ValidateFunc) (Planfile, error) {
	var unsafePlanfile UnsafePlanFile
	if err := yaml.NewDecoder(reader).Decode(&unsafePlanfile); err != nil {
		return Planfile{}, err
	}
	unsafePlanfile.Path = planfilePath

	planFile, validationErr := validate(unsafePlanfile)
	if validationErr != nil {
		return Planfile{}, validationErr
	}

	return planFile, nil
}
