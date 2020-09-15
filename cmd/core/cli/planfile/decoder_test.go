package planfile

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		planfilePath FilePath
		yaml         string
		expected     Planfile
	}{
		{
			planfilePath: "/path/to/planfile.yml",
			yaml:         `instance_groups: {}`,
			expected:     Planfile{},
		},
		{
			planfilePath: "/path/to/planfile.yml",
			yaml: `
instance_groups:
  group1:
    - platform: any-platform
      ios: 12.0
      device: apple iphone xs
      ipa: "./app.ipa"
      lifetime_sec: 100
`,
			expected: Planfile{
				plans: []platforms.EitherPlan{
					platforms.NewIOSPlan(
						"any-platform",
						"group1",
						platforms.NewIOSDevice(
							"apple iphone xs",
							"12.0",
						),
						"/path/to/app.ipa",
						[]string{},
						100*time.Second,
						`at 1-th plan for instance group "group1"`,
					).Either(),
				},
			},
		},
		{
			planfilePath: "/path/to/planfile.yml",
			yaml: `
instance_groups:
  group2:
    - platform: any-platform
      android: 9
      device: google google pixel3
      apk: ./app.apk
      app_id: com.example.app
      lifetime_sec: 200
`,
			expected: Planfile{
				plans: []platforms.EitherPlan{
					platforms.NewAndroidPlan(
						"any-platform",
						"group2",
						platforms.NewAndroidDevice(
							"google google pixel3",
							"9",
						),
						"/path/to/app.apk",
						"com.example.app",
						[]string{},
						200*time.Second,
						`at 1-th plan for instance group "group2"`,
					).Either(),
				},
			},
		},
		{
			planfilePath: "/path/to/planfile.yml",
			yaml: `
instance_groups:
  group1:
    - platform: any-platform
      ios: 12.0
      device: apple iphone xs
      ipa: ./app.ipa
      args:
        - -ARG1
        - VALUE1
      lifetime_sec: 100
  group2:
    - platform: any-platform
      android: 9
      device: google google pixel3
      apk: ./app.apk
      app_id: com.example.app
      intent_extras:
        - -e
        - ARG1
        - VALUE1
      lifetime_sec: 200
`,
			expected: Planfile{
				plans: []platforms.EitherPlan{
					platforms.NewIOSPlan(
						"any-platform",
						"group1",
						platforms.NewIOSDevice(
							"apple iphone xs",
							"12.0",
						),
						"/path/to/app.ipa",
						[]string{"-ARG1", "VALUE1"},
						100*time.Second,
						`at 1-th plan for instance group "group1"`,
					).Either(),
					platforms.NewAndroidPlan(
						"any-platform",
						"group2",
						platforms.NewAndroidDevice(
							"google google pixel3",
							"9",
						),
						"/path/to/app.apk",
						"com.example.app",
						[]string{"-e", "ARG1", "VALUE1"},
						200*time.Second,
						`at 1-th plan for instance group "group2"`,
					).Either(),
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Decode(strings.NewReader(%q))", c.yaml), func(t *testing.T) {
			got, err := Decode(c.planfilePath, strings.NewReader(c.yaml), NewValidateFunc(exec.AnySuccessfulStatFunc()))

			if err != nil {
				t.Errorf("got (_, %v), want (_, nil)", err)
				return
			}

			if !reflect.DeepEqual(got.Plans(), c.expected.Plans()) {
				t.Error(cmp.Diff(c.expected.Plans(), got.Plans()))
				return
			}
		})
	}
}
