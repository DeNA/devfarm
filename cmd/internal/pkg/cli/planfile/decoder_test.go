package planfile

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		yaml          string
		expected      Planfile
		expectedError bool
	}{
		{
			yaml:          `instance_groups: {}`,
			expected:      Planfile{},
			expectedError: false,
		},
		{
			yaml: `
instance_groups:
  group1:
    - platform: any-platform
      ios: 12.0
      device: apple iphone xs
      ipa: "path/to/app.ipa"
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
						"path/to/app.ipa",
						[]string{},
						100*time.Second,
						`at 1-th plan for instance group "group1"`,
					).Either(),
				},
			},
			expectedError: false,
		},
		{
			yaml: `
instance_groups:
  group2:
    - platform: any-platform
      android: 9
      device: google google pixel3
      apk: path/to/app.apk
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
						"path/to/app.apk",
						"com.example.app",
						[]string{},
						200*time.Second,
						`at 1-th plan for instance group "group2"`,
					).Either(),
				},
			},
			expectedError: false,
		},
		{
			yaml: `
instance_groups:
  group1:
    - platform: any-platform
      ios: 12.0
      device: apple iphone xs
      ipa: "path/to/app.ipa"
      args:
        - -ARG1
        - VALUE1
      lifetime_sec: 100
  group2:
    - platform: any-platform
      android: 9
      device: google google pixel3
      apk: path/to/app.apk
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
						"path/to/app.ipa",
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
						"path/to/app.apk",
						"com.example.app",
						[]string{"-e", "ARG1", "VALUE1"},
						200*time.Second,
						`at 1-th plan for instance group "group2"`,
					).Either(),
				},
			},
			expectedError: false,
		},
		{
			yaml:          ``,
			expected:      Planfile{},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Decode(strings.NewReader(%q))", c.yaml), func(t *testing.T) {
			got, err := Decode(strings.NewReader(c.yaml))

			if c.expectedError {
				if err == nil {
					t.Errorf("got (_, nil), want (_, error)")
					return
				}
			} else {
				if err != nil {
					t.Errorf("got (_, %v), want (_, nil)", err)
					return
				}

				if !reflect.DeepEqual(got.Plans(), c.expected.Plans()) {
					t.Error(cmp.Diff(c.expected.Plans(), got.Plans()))
					return
				}
			}
		})
	}
}
