package devicefarm

func StubProjectLister(projects []Project, err error) ProjectLister {
	return func() ([]Project, error) {
		return projects, err
	}
}
