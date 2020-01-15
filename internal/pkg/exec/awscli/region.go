package awscli

type Region string

const (
	RegionIsUSWest2 Region = "us-west-2"
)

func (region Region) Name() string {
	return string(region)
}
