package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestFindProjectARN(t *testing.T) {
	var projectARN devicefarm.ProjectARN = "arn:devicefarm:EXAMPLE"
	var anotherProjectARN devicefarm.ProjectARN = "arn:devicefarm:ANOTHER"

	cases := []struct {
		groupName        platforms.InstanceGroupName
		projects         []devicefarm.Project
		projectsError    error
		expected         devicefarm.ProjectARN
		expectedError    bool
		expectedNotFound bool
	}{
		{
			groupName: "example",
			projects: []devicefarm.Project{
				devicefarm.NewProject(
					"devfarm-example",
					projectARN,
				),
				devicefarm.NewProject(
					"devfarm-another",
					anotherProjectARN,
				),
			},
			projectsError:    nil,
			expected:         projectARN,
			expectedError:    false,
			expectedNotFound: false,
		},
		{
			groupName:        "example",
			projects:         []devicefarm.Project{},
			projectsError:    nil,
			expected:         "",
			expectedError:    true,
			expectedNotFound: true,
		},
		{
			groupName:        "example",
			projects:         []devicefarm.Project{},
			projectsError:    testutil.AnyError,
			expected:         "",
			expectedError:    true,
			expectedNotFound: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("FindProjectARN(%v, client); client.ListProjects() == (%v, %v)", c.groupName, c.projects, c.projectsError), func(t *testing.T) {
			logger := logging.NullSeverityLogger()
			listProjects := devicefarm.StubProjectLister(c.projects, c.projectsError)
			findProjectARN := newProjectARNFinder(logger, listProjects)

			got, err := findProjectARN(c.groupName)

			if c.expectedError {
				if err == nil {
					t.Errorf("got (_, nil), want (_, error)")
				} else if c.expectedNotFound && err.NotFound == nil {
					t.Errorf("got (_, &projectARNFinderError{NotFound:nil}), want (_, &projectARNFinderError{NotFound:error})")
				}
			} else {
				if got != c.expected {
					t.Errorf("got (%v, nil), want (%v, nil)", got, c.expected)
				}
			}
		})
	}
}
