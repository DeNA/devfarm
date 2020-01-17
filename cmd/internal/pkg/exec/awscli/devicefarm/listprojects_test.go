package devicefarm

import (
	"reflect"
	"testing"
)

func TestNewProjectLister(t *testing.T) {
	execute := StubExecutor([]byte(listProjectsJSONExample), []byte{}, nil)
	listProjects := NewProjectLister(execute)

	projects, err := listProjects()

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if !reflect.DeepEqual(projects, listProjectsExample) {
		t.Errorf("got (%v, nil), want (%v, nil)", projects, listProjectsExample)
		return
	}
}
