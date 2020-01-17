package planfile

import (
	"github.com/dena/devfarm/cmd/core/platforms"
	"sort"
	"time"
)

type FilePath string

type Planfile struct {
	plans []platforms.EitherPlan
	dirty bool
}

func NewPlanfile(plans ...platforms.EitherPlan) *Planfile {
	return &Planfile{plans: plans, dirty: true}
}

func NewTemplate() *Planfile {
	return &Planfile{plans: []platforms.EitherPlan{
		platforms.NewIOSPlan(
			"aws-device-farm",
			"example",
			platforms.IOSDevice{
				DeviceName: "apple iphone xs",
				OSVersion:  "12.0",
			},
			"path/to/your.ipa",
			platforms.IOSArgs{},
			15*time.Minute,
			platforms.LocationHintForGeneratedByCode,
		).Either(),
		platforms.NewAndroidPlan(
			"aws-device-farm",
			"example",
			platforms.AndroidDevice{
				DeviceName: "google google pixel 3",
				OSVersion:  "9.0",
			},
			"path/to/your.apk",
			"com.example.YourApp",
			platforms.AndroidIntentExtras{},
			15*time.Minute,
			platforms.LocationHintForGeneratedByCode,
		).Either(),
	}}
}

func (planfile *Planfile) AddPlan(plan platforms.EitherPlan) {
	planfile.dirty = true
	planfile.plans = append(planfile.plans, plan)
}

func (planfile *Planfile) Plans() []platforms.EitherPlan {
	// FIXME: Prevent unnecessary sorting
	sort.Slice(planfile.plans, func(i, j int) bool {
		if planfile.plans[i].LocationHint != planfile.plans[j].LocationHint {
			return planfile.plans[i].LocationHint < planfile.plans[j].LocationHint
		}

		if planfile.plans[i].GroupName != planfile.plans[j].GroupName {
			return planfile.plans[i].GroupName < planfile.plans[j].GroupName
		}

		if planfile.plans[i].Platform != planfile.plans[j].Platform {
			return planfile.plans[i].Platform < planfile.plans[j].Platform
		}

		if planfile.plans[i].OSName != planfile.plans[j].OSName {
			return planfile.plans[i].OSName < planfile.plans[j].OSName
		}

		return planfile.plans[i].IOSSpecificPart.Device.Less(planfile.plans[j].IOSSpecificPart.Device)
	})
	return planfile.plans
}
