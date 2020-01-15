package platforms

type AllInstanceLister func() ([]InstanceOrError, error)
