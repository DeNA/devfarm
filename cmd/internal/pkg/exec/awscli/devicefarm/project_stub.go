package devicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli"
)

// NOTE: https://docs.aws.amazon.com/cli/latest/reference/devicefarm/create-project.html
var createProjectResponseJSONExample = fmt.Sprintf(`{ "project": %s }`, projectJSONExample)

var listProjectsJSONExample = fmt.Sprintf(`{ "projects": [%s] }`, projectJSONExample)

var listProjectsExample = []Project{ProjectExample}

var projectJSONExample = `{
	"name": "myproject",
	"arn": "arn:aws:devicefarm:us-west-2:123456789012:project:070fc3ca-7ec1-4741-9c1f-d3e044efc506",
	"created": 1503612890.057
}`

var ProjectExample = NewProject(
	"myproject",
	"arn:aws:devicefarm:us-west-2:123456789012:project:070fc3ca-7ec1-4741-9c1f-d3e044efc506",
	awscli.NewTimestamp(1503612890),
)

func AnyProject() Project {
	return NewProject(
		"ANY_PROJECT",
		"arn:devicefarm:ANY_ARN",
		awscli.NewTimestamp(0),
	)
}
