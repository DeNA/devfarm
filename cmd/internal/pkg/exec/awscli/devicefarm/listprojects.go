package devicefarm

import "encoding/json"

type ProjectLister func() ([]Project, error)

type listProjectsResponse struct {
	Projects []Project `json:"projects"`
}

func NewProjectLister(deviceFarmCmd Executor) ProjectLister {
	return func() ([]Project, error) {
		// NOTE: https://docs.aws.amazon.com/cli/latest/reference/devicefarm/list-projects.html
		result, execErr := deviceFarmCmd(
			"list-projects",
		)
		if execErr != nil {
			return nil, execErr
		}

		var response listProjectsResponse
		if err := json.Unmarshal(result.Stdout, &response); err != nil {
			return nil, err
		}

		return response.Projects, nil
	}
}
