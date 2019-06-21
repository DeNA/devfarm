package cli

import "testing"

func TestOnlyCommonOptions(t *testing.T) {
	type VerboseAndDryRun struct {
		Verbose bool
		DryRun  bool
	}

	cases := []struct {
		args     []string
		expected *VerboseAndDryRun
	}{
		{
			args: []string{},
			expected: &VerboseAndDryRun{
				Verbose: false,
				DryRun:  false,
			},
		},
		{
			args: []string{"--verbose"},
			expected: &VerboseAndDryRun{
				Verbose: true,
				DryRun:  false,
			},
		},
		{
			args: []string{"--dry-run"},
			expected: &VerboseAndDryRun{
				Verbose: false,
				DryRun:  true,
			},
		},
		{
			args: []string{"--verbose", "--dry-run"},
			expected: &VerboseAndDryRun{
				Verbose: true,
				DryRun:  true,
			},
		},
		{
			args:     []string{"--help"},
			expected: nil,
		},
		{
			args:     []string{"--help", "--verbose"},
			expected: nil,
		},
		{
			args:     []string{"--something-wrong"},
			expected: nil,
		},
	}

	for _, c := range cases {
		got, err := TakeOnlyVerboseAndDryRunOpts(c.args)

		if c.expected == nil {
			if got != nil {
				t.Errorf("TakeOnlyVerboseAndDryRunOpts(%v) == %v, but wanted %v (err=%v)", c.args, got, c.expected, err)
			}
		} else {
			if got == nil {
				t.Errorf("TakeOnlyVerboseAndDryRunOpts(%v) == nil, but wanted %v (err=%v)", c.args, c.expected, err)
			} else if *got != *c.expected {
				t.Errorf("TakeOnlyVerboseAndDryRunOpts(%v) == %v, but wanted %v (err=%v)", c.args, got, c.expected, err)
			}
		}
	}
}
