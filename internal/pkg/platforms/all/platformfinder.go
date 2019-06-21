package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func FindPlatformByName(platformName string) (platforms.Platform, error) {
	for _, platform := range PlatformTable {
		if string(platform.ID()) == platformName {
			return platform, nil
		}
	}

	return nil, fmt.Errorf("no such platform: %q (availables are %v)", platformName, PlatformNames())
}
