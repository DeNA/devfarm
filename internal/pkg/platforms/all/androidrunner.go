package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func RunAndroid(
	plan platforms.AndroidPlan,
	bag platforms.AndroidForeverBag,
) error {
	platform, ok := PlatformTable[plan.CommonPart.Platform]
	if !ok {
		return fmt.Errorf("no such platform: %q", plan.CommonPart.Platform)
	}

	runAndroid := platform.AndroidRunner()
	return runAndroid(plan, bag)
}
