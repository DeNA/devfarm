package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

type instanceGroupLister func() ([]platforms.InstanceGroupListEntry, error)

func newInstanceGroupLister(listProjects devicefarm.ProjectLister) instanceGroupLister {
	return func() ([]platforms.InstanceGroupListEntry, error) {
		projects, projectsErr := listProjects()
		if projectsErr != nil {
			return nil, projectsErr
		}

		return mapProjectsToInstanceGroups(projects), nil
	}
}

func mapProjectsToInstanceGroups(projects []devicefarm.Project) []platforms.InstanceGroupListEntry {
	entries := make([]platforms.InstanceGroupListEntry, 0)

	for _, project := range projects {
		groupName, err := project.Name.ToInstanceGroupName()

		if err != nil {
			if err.Unmanaged != nil {
				// NOTE: Ignore not managed projects.
			} else {
				group := platforms.NewErrorInstanceGroup()
				entry := platforms.NewInstanceGroupListEntry(group, err.Unspecified)
				entries = append(entries, entry)
			}
		} else {
			group := platforms.NewInstanceGroup(groupName)
			entry := platforms.NewInstanceGroupListEntry(group, nil)
			entries = append(entries, entry)
		}
	}

	return entries
}
