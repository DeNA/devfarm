package remoteagent

import (
	"github.com/dena/devfarm/cmd/core/platforms"
)

type AppPathOnRemote string

func (p AppPathOnRemote) OSName() (platforms.OSName, error) {
	return platforms.DetectOSName(string(p))
}
