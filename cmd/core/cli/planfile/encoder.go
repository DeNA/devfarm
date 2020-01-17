package planfile

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Encode(planfile Planfile, writer io.Writer) error {
	encoder := yaml.NewEncoder(writer)
	if err := encoder.Encode(NewUnsafePlanFile(planfile)); err != nil {
		return err
	}
	return nil
}
