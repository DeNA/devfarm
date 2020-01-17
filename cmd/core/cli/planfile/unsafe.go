package planfile

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"math"
)

type UnsafePlanFile struct {
	InstanceGroups map[UnsafeInstanceGroupName][]UnsafePlan `yaml:"instance_groups"`
}

type UnsafeInstanceGroupName string

func NewUnsafePlanFile(planfile Planfile) *UnsafePlanFile {
	unsafePlanfile := &UnsafePlanFile{}
	unsafePlanfile.InstanceGroups = make(map[UnsafeInstanceGroupName][]UnsafePlan)
	for _, plan := range planfile.plans {
		var unsafePlan UnsafePlan
		var unsafeGroupName UnsafeInstanceGroupName

		switch plan.OSName {
		case platforms.OSIsIOS:
			iosPlan := plan.IOS()
			unsafePlan = UnsafePlan{
				Platform:    string(iosPlan.CommonPart.Platform),
				Device:      string(iosPlan.IOSSpecificPart.Device.DeviceName),
				IOS:         string(iosPlan.IOSSpecificPart.Device.OSVersion),
				IPA:         string(iosPlan.IOSSpecificPart.IPA),
				IOSArgs:     iosPlan.IOSSpecificPart.Args,
				LifetimeSec: int(math.Ceil(iosPlan.CommonPart.Lifetime.Seconds())),
			}
			unsafeGroupName = UnsafeInstanceGroupName(iosPlan.CommonPart.GroupName)

		case platforms.OSIsAndroid:
			androidPlan := plan.Android()
			unsafePlan = UnsafePlan{
				Platform:     string(androidPlan.CommonPart.Platform),
				Device:       string(androidPlan.AndroidSpecificPart.Device.DeviceName),
				Android:      string(androidPlan.AndroidSpecificPart.Device.OSVersion),
				APK:          string(androidPlan.AndroidSpecificPart.APK),
				IntentExtras: androidPlan.AndroidSpecificPart.IntentExtras,
				AndroidAppID: string(androidPlan.AndroidSpecificPart.AppID),
				LifetimeSec:  int(math.Ceil(androidPlan.CommonPart.Lifetime.Seconds())),
			}
			unsafeGroupName = UnsafeInstanceGroupName(androidPlan.CommonPart.GroupName)

		default:
			panic(fmt.Sprintf("unsupported os: %q", plan.OSName))
		}

		unsafePlans, ok := unsafePlanfile.InstanceGroups[unsafeGroupName]
		if ok {
			unsafePlanfile.InstanceGroups[unsafeGroupName] = append(unsafePlans, unsafePlan)
		} else {
			unsafePlanfile.InstanceGroups[unsafeGroupName] = []UnsafePlan{unsafePlan}
		}
	}
	return unsafePlanfile
}

type UnsafePlan struct {
	Platform    string `yaml:"platform"`
	Device      string `yaml:"device"`
	LifetimeSec int    `yaml:"lifetime_sec"`

	IOS     string   `yaml:"ios,omitempty"`
	IPA     string   `yaml:"ipa,omitempty"`
	IOSArgs []string `yaml:"args,omitempty"`

	Android      string   `yaml:"android,omitempty"`
	APK          string   `yaml:"apk,omitempty"`
	IntentExtras []string `yaml:"intent_extras,omitempty"`
	AndroidAppID string   `yaml:"app_id,omitempty"`
}

func (part UnsafePlan) DetectOS() platforms.OSName {
	if len(part.IOS) > 0 {
		return platforms.OSIsIOS
	}

	if len(part.Android) > 0 {
		return platforms.OSIsAndroid
	}

	return platforms.OSIsUnavailable
}
