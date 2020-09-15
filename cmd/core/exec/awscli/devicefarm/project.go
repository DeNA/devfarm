package devicefarm

import (
	"encoding/json"
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"strings"
)

type ProjectARN string

func (p ProjectARN) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

func (p *ProjectARN) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*p = ProjectARN(s)
	return nil
}

type Project struct {
	Name    ProjectName      `json:"name"`
	ARN     ProjectARN       `json:"arn"`
}

func NewProject(name ProjectName, arn ProjectARN) Project {
	return Project{
		Name:    name,
		ARN:     arn,
	}
}

type ProjectName string

var projectNamePrefix = "devfarm-"

type InstanceGroupNameError struct {
	Unmanaged   error
	Unspecified error
}

func (e InstanceGroupNameError) Error() string {
	if e.Unmanaged != nil {
		return e.Unmanaged.Error()
	}
	return e.Unspecified.Error()
}

func FromInstanceGroupName(name platforms.InstanceGroupName) ProjectName {
	return ProjectName(fmt.Sprintf("%s%s", projectNamePrefix, string(name)))
}

func (p ProjectName) ToInstanceGroupName() (platforms.InstanceGroupName, *InstanceGroupNameError) {
	s := string(p)

	if !strings.HasPrefix(s, projectNamePrefix) {
		return "", &InstanceGroupNameError{
			Unmanaged:   fmt.Errorf("not devfarm-managed project: %q", s),
			Unspecified: nil,
		}
	}

	nameText := s[len(projectNamePrefix):]
	name, nameErr := platforms.NewInstanceGroupName(nameText)

	if nameErr != nil {
		return "", &InstanceGroupNameError{
			Unmanaged:   nil,
			Unspecified: nameErr,
		}
	}

	return name, nil
}

func (p ProjectName) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

func (p *ProjectName) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*p = ProjectName(s)
	return nil
}
