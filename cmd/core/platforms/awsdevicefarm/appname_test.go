package awsdevicefarm

import (
	"fmt"
	"testing"
)

func TestValidateAppName(t *testing.T) {
	cases := []struct {
		appPath         ipaOrApkPathOnLocal
		expectedAppName appName
		expectedExtname string
		expectedErr     bool
	}{
		{
			appPath:         "path/to/app.ipa",
			expectedAppName: "app",
			expectedExtname: ".ipa",
			expectedErr:     false,
		},
		{
			appPath:         "path/to/app.apk",
			expectedAppName: "app",
			expectedExtname: ".apk",
			expectedErr:     false,
		},
		{
			appPath:         "path/to/マルチバイト.apk",
			expectedAppName: "",
			expectedExtname: "",
			expectedErr:     true,
		},
		{
			appPath:         "",
			expectedAppName: "",
			expectedExtname: "",
			expectedErr:     true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("validateAppName(%q)", c.appPath), func(t *testing.T) {
			got, extname, err := validateAppName(c.appPath)

			if c.expectedErr {
				if err == nil {
					t.Errorf("got (_, _, nil), want (_, _, error)")
				}
			} else {
				if err != nil {
					t.Errorf("got (_, _, %v), want (_, _, nil)", err)
					return
				}

				if got != c.expectedAppName {
					t.Errorf("got (%v, _, nil), want (%v, _, nil)", got, c.expectedAppName)
				}

				if extname != c.expectedExtname {
					t.Errorf("got (_, %v, nil), want (_, %v, nil)", extname, c.expectedExtname)
				}
			}
		})
	}
}
