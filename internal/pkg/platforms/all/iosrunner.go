package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func RunIOS(
	plan platforms.IOSPlan,
	bag platforms.IOSForeverBag,
) error {
	platform, ok := PlatformTable[plan.CommonPart.Platform]
	if !ok {
		return fmt.Errorf("no such platform: %q", plan.CommonPart.Platform)
	}

	runIOS := platform.IOSRunner()
	return runIOS(plan, bag)
}
