package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func ForeverIOS(
	plan platforms.IOSPlan,
	bag platforms.IOSForeverBag,
) error {
	platform, ok := PlatformTable[plan.CommonPart.Platform]
	if !ok {
		return fmt.Errorf("no such platform: %q", plan.CommonPart.Platform)
	}

	foreverIOS := platform.IOSForever()
	return foreverIOS(plan, bag)
}
