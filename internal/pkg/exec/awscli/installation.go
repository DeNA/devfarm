package awscli

import "github.com/dena/devfarm/internal/pkg/exec"

type InstalltionStatusGetter func() error

func NewInstallStatusGetter(find exec.ExecutableFinder) InstalltionStatusGetter {
	return func() error {
		return find("aws")
	}
}
