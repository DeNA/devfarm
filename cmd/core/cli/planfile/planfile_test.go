package planfile

import (
	"encoding/json"
	"github.com/dena/devfarm/cmd/core/platforms"
	"testing"
	"time"
)

func TestJSON(t *testing.T) {
	planFile := Planfile{
		plans: []platforms.EitherPlan{
			platforms.NewIOSPlan(
				"any-platform",
				"group1",
				platforms.NewIOSDevice(
					"apple iphone xs",
					"12.0",
				),
				"path/to/app.ipa",
				[]string{"-ARG1", "VALUE1"},
				time.Minute,
				"any location",
			).Either(),
			platforms.NewAndroidPlan(
				"any-platform",
				"group2",
				platforms.NewAndroidDevice(
					"google google pixel3",
					"9",
				),
				"path/to/app.apk",
				"com.example.app",
				[]string{"-e", "ARG1", "VALUE1"},
				time.Minute,
				"any location",
			).Either(),
		},
	}

	if _, err := json.Marshal(planFile.Plans()); err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
	}
}
