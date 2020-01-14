package devicefarm

import (
	"encoding/json"
	"fmt"
	"time"
)

type ProjectCreator func(name ProjectName, timeout time.Duration) (Project, error)

type createProjectResponse struct {
	Project Project `json:"project"`
}

func NewProjectCreator(deviceFarmCmd Executor) ProjectCreator {
	return func(name ProjectName, timeout time.Duration) (Project, error) {
		// NOTE: https://docs.aws.amazon.com/cli/latest/reference/devicefarm/create-project.html
		result, execErr := deviceFarmCmd(
			"create-project",
			"--name", string(name),
			"--default-job-timeout-minutes", fmt.Sprintf("%d", int(timeout.Minutes())),
		)
		if execErr != nil {
			return Project{}, execErr
		}

		var response createProjectResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return Project{}, err
		}

		return response.Project, nil
	}
}
