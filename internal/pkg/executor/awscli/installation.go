package awscli

import "github.com/dena/devfarm/internal/pkg/executor"

type InstalltionStatusGetter func() error

func NewInstallStatusGetter(find executor.ExecutableFinder) InstalltionStatusGetter {
	return func() error {
		return find("aws")
	}
}
