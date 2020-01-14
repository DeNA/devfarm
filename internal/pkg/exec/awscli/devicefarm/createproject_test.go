package devicefarm

import "testing"

func TestNewProjectCreator(t *testing.T) {
	execute := StubExecutor([]byte(createProjectResponseJSONExample), []byte{}, nil)
	createProject := NewProjectCreator(execute)

	_, err := createProject(ProjectName("example"), 0)

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}
}
