package planfile

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"reflect"
	"testing"
	"time"
)

func TestUnsafePlanfileIOS(t *testing.T) {
	groupName := "ANY_INSTANCE_GROUP"

	cases := []struct {
		unsafePlan    UnsafePlan
		locationHint  string
		expected      platforms.IOSPlan
		expectedError bool
	}{
		{
			unsafePlan: UnsafePlan{
				Platform:    "any platform",
				Device:      "apple iphone xs",
				IOS:         "12",
				IPA:         "path/to/app.ipa",
				IOSArgs:     []string{"-ARG1", "VALUE"},
				LifetimeSec: 100,
			},
			locationHint: "location-1",
			expected: platforms.NewIOSPlan(
				"any platform",
				platforms.InstanceGroupName(groupName),
				platforms.NewIOSDevice(
					"apple iphone xs",
					"12",
				),
				"path/to/app.ipa",
				[]string{"-ARG1", "VALUE"},
				100*time.Second,
				"location-1",
			),
			expectedError: false,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:    "any platform",
				Device:      "apple iphone xs",
				IOS:         "12",
				IPA:         "path/to/app.ipa",
				LifetimeSec: 200,
			},
			locationHint: "location-2",
			expected: platforms.NewIOSPlan(
				"any platform",
				platforms.InstanceGroupName(groupName),
				platforms.NewIOSDevice(
					"apple iphone xs",
					"12",
				),
				"path/to/app.ipa",
				[]string{},
				200*time.Second,
				"location-2",
			),
			expectedError: false,
		},
		{
			unsafePlan: UnsafePlan{
				Device:      "apple iphone xs",
				IOS:         "12",
				IPA:         "path/to/app.ipa",
				LifetimeSec: 300,
			},
			locationHint:  "location-3",
			expected:      platforms.IOSPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:    "any platform",
				IOS:         "12",
				IPA:         "path/to/app.ipa",
				LifetimeSec: 300,
			},
			locationHint:  "location-3",
			expected:      platforms.IOSPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:    "any platform",
				Device:      "apple iphone xs",
				IOS:         "12",
				LifetimeSec: 300,
			},
			locationHint:  "location-3",
			expected:      platforms.IOSPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform: "any platform",
				Device:   "apple iphone xs",
				IOS:      "12",
				IPA:      "path/to/app.ipa",
			},
			locationHint:  "location-3",
			expected:      platforms.IOSPlan{},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v.DetectOS()", c.unsafePlan), func(t *testing.T) {
			got := c.unsafePlan.DetectOS()

			if got != platforms.OSIsIOS {
				t.Errorf("got (%v, nil), want (%v, nil)", got, platforms.OSIsIOS)
				return
			}
		})

		t.Run(fmt.Sprintf("%v.ValidateAsIOS(_)", c.unsafePlan), func(t *testing.T) {
			got, err := ValidateAsIOS(c.unsafePlan, platforms.InstanceGroupName(groupName), c.locationHint)

			if c.expectedError {
				if err == nil {
					t.Errorf("got nil, want error")
					return
				}
			} else {
				if err != nil {
					t.Errorf("got (_, %v), want (_, nil)", err)
					return
				}

				if !reflect.DeepEqual(got, c.expected) {
					t.Error(cmp.Diff(c.expected, got))
					return
				}
			}
		})
	}
}

func TestRawPlanFileAndroid(t *testing.T) {
	groupName := "ANY_INSTANCE_GROUP"

	cases := []struct {
		unsafePlan    UnsafePlan
		locationHint  string
		expected      platforms.AndroidPlan
		expectedError bool
	}{
		{
			unsafePlan: UnsafePlan{
				Platform:     "any platform",
				Device:       "google google pixel3",
				Android:      "9",
				AndroidAppID: "com.example.app",
				APK:          "path/to/app.apk",
				IntentExtras: []string{"-e", "ARG1", "VALUE"},
				LifetimeSec:  100,
			},
			locationHint: "location-1",
			expected: platforms.NewAndroidPlan(
				"any platform",
				platforms.InstanceGroupName(groupName),
				platforms.NewAndroidDevice(
					"google google pixel3",
					"9",
				),
				"path/to/app.apk",
				"com.example.app",
				[]string{"-e", "ARG1", "VALUE"},
				100*time.Second,
				"location-1",
			),
			expectedError: false,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:     "any platform",
				Device:       "google google pixel3",
				Android:      "9",
				AndroidAppID: "com.example.app",
				APK:          "path/to/app.apk",
				LifetimeSec:  200,
			},
			locationHint: "location-2",
			expected: platforms.NewAndroidPlan(
				"any platform",
				platforms.InstanceGroupName(groupName),
				platforms.NewAndroidDevice(
					"google google pixel3",
					"9",
				),
				"path/to/app.apk",
				"com.example.app",
				[]string{},
				200*time.Second,
				"location-2",
			),
			expectedError: false,
		},
		{
			unsafePlan: UnsafePlan{
				Device:       "google google pixel3",
				Android:      "9",
				APK:          "path/to/app.apk",
				AndroidAppID: "com.example.app",
				LifetimeSec:  300,
			},
			locationHint:  "location-3",
			expected:      platforms.AndroidPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:     "any platform",
				Android:      "9",
				APK:          "path/to/app.apk",
				AndroidAppID: "com.example.app",
				LifetimeSec:  300,
			},
			locationHint:  "location-3",
			expected:      platforms.AndroidPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:     "any platform",
				Device:       "google google pixel3",
				Android:      "9",
				AndroidAppID: "com.example.app",
				LifetimeSec:  300,
			},
			locationHint:  "location-3",
			expected:      platforms.AndroidPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:    "any platform",
				Device:      "google google pixel3",
				Android:     "9",
				APK:         "path/to/app.apk",
				LifetimeSec: 300,
			},
			locationHint:  "location-3",
			expected:      platforms.AndroidPlan{},
			expectedError: true,
		},
		{
			unsafePlan: UnsafePlan{
				Platform:     "any platform",
				Device:       "google google pixel3",
				Android:      "9",
				APK:          "path/to/app.apk",
				AndroidAppID: "com.example.app",
			},
			locationHint:  "location-3",
			expected:      platforms.AndroidPlan{},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v.DetectOS()", c.unsafePlan), func(t *testing.T) {
			got := c.unsafePlan.DetectOS()

			if got != platforms.OSIsAndroid {
				t.Errorf("got (%v, nil), want (%v, nil)", got, platforms.OSIsAndroid)
				return
			}
		})

		t.Run(fmt.Sprintf("%v.ValidateAsAndroid(_)", c.unsafePlan), func(t *testing.T) {
			got, err := ValidateAsAndroid(c.unsafePlan, platforms.InstanceGroupName(groupName), c.locationHint)

			if c.expectedError {
				if err == nil {
					t.Errorf("got nil, want error")
					return
				}
			} else {
				if err != nil {
					t.Errorf("got (_, %v), want (_, nil)", err)
					return
				}

				if !reflect.DeepEqual(got, c.expected) {
					t.Error(cmp.Diff(c.expected, got))
					return
				}
			}
		})
	}
}

func TestNewUnsafePlanFile(t *testing.T) {
	cases := []struct {
		desc     string
		planfile Planfile
		expected UnsafePlanFile
	}{
		{
			desc: "ios",
			planfile: Planfile{plans: []platforms.EitherPlan{
				platforms.NewIOSPlan(
					"platform-1",
					"example-group-1",
					platforms.IOSDevice{
						DeviceName: "apple iphone xs",
						OSVersion:  "12.0",
					},
					"path/to/app.ipa",
					platforms.IOSArgs{"-ARG"},
					1*time.Second,
					"location hint will be lost",
				).Either(),
			}},
			expected: UnsafePlanFile{
				InstanceGroups: map[UnsafeInstanceGroupName][]UnsafePlan{
					"example-group-1": {
						{
							Platform:    "platform-1",
							Device:      "apple iphone xs",
							IOS:         "12.0",
							IPA:         "path/to/app.ipa",
							IOSArgs:     platforms.IOSArgs{"-ARG"},
							LifetimeSec: 1,
						},
					},
				},
			},
		},
		{
			desc: "android",
			planfile: Planfile{plans: []platforms.EitherPlan{
				platforms.NewAndroidPlan(
					"platform-2",
					"example-group-2",
					platforms.AndroidDevice{
						DeviceName: "google google pixel 3",
						OSVersion:  "9.0",
					},
					"path/to/app.apk",
					"com.example.apk",
					platforms.AndroidIntentExtras{"-e", "ARG"},
					2*time.Second,
					"location hint will be lost",
				).Either(),
			}},
			expected: UnsafePlanFile{
				InstanceGroups: map[UnsafeInstanceGroupName][]UnsafePlan{
					"example-group-2": {
						{
							Platform:     "platform-2",
							Device:       "google google pixel 3",
							Android:      "9.0",
							APK:          "path/to/app.apk",
							IntentExtras: platforms.AndroidIntentExtras{"-e", "ARG"},
							AndroidAppID: "com.example.apk",
							LifetimeSec:  2,
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := NewUnsafePlanFile(c.planfile)

			if !reflect.DeepEqual(*got, c.expected) {
				t.Error(cmp.Diff(c.expected, *got))
				return
			}
		})
	}
}
