package adb

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/testutil"
	"testing"
)

func TestNewMainIntentFinder(t *testing.T) {
	cases := []struct {
		stdoutBytes   []byte
		adbError      *ExecutorError
		expected      Intent
		expectedError bool
	}{
		{
			stdoutBytes: []byte(`
DUMP OF SERVICE package:
  Activity Resolver Table:
    Non-Data Actions:
        android.intent.action.MAIN:
          aba9035 com.DeNA.GodlikeReborn/com.unity3d.player.UnityPlayerActivity filter 8416d53
            Action: "android.intent.action.MAIN"
            Category: "android.intent.category.LAUNCHER"
            Category: "android.intent.category.LEANBACK_LAUNCHER"

  Key Set Manager:
    [com.DeNA.GodlikeReborn]
        Signing KeySets: 5
`),
			adbError:      nil,
			expected:      "com.DeNA.GodlikeReborn/com.unity3d.player.UnityPlayerActivity",
			expectedError: false,
		},
		{
			stdoutBytes:   []byte{},
			adbError:      nil,
			expectedError: true,
		},
		{
			stdoutBytes:   []byte{},
			adbError:      &ExecutorError{NoSuchCommand: testutil.AnyError},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("adb.Executor(_) == ([]byte(%q), %v)", string(c.stdoutBytes), c.adbError), func(t *testing.T) {
			execute := StubExecutor(c.stdoutBytes, c.adbError)
			findMainIntent := NewMainIntentFinder(execute)

			got, err := findMainIntent("ANY-SERIAL-NUMBER", "com.DeNA.GodlikeReborn")
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

				if got != c.expected {
					t.Errorf("got (%v, nil), want (%v, nil)", got, c.expected)
					return
				}
			}
		})
	}
}
