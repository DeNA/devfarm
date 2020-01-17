package awscli

import "github.com/dena/devfarm/cmd/core/exec"

type InstalltionStatusGetter func() error

func NewInstallStatusGetter(find exec.ExecutableFinder) InstalltionStatusGetter {
	return func() error {
		return find("aws")
	}
}
