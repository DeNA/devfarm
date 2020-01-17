package all

import (
	"errors"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestResultTable_TextTable(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	header := []string{"platform", "status", "note"}

	cases := []struct {
		desc     string
		table    ResultTable
		expected [][]string
	}{
		{
			desc:  "empty table",
			table: ResultTable{},
			expected: [][]string{
				header,
			},
		},
		{
			desc: "empty results",
			table: ResultTable{
				"platform-1": *platforms.NewResults(),
			},
			expected: [][]string{
				header,
			},
		},
		{
			desc: "1 platform, 1 success",
			table: ResultTable{
				"platform-1": *platforms.NewResults(nil),
			},
			expected: [][]string{
				header,
				{"platform-1", "ok", ""},
			},
		},
		{
			desc: "1 platform, 2 success",
			table: ResultTable{
				"platform-1": *platforms.NewResults(nil, nil),
			},
			expected: [][]string{
				header,
				{"platform-1", "ok", ""},
				{"platform-1", "ok", ""},
			},
		},
		{
			desc: "1 platform, 1 error",
			table: ResultTable{
				"platform-1": *platforms.NewResults(err1),
			},
			expected: [][]string{
				header,
				{"platform-1", "error", "err1"},
			},
		},
		{
			desc: "1 platform, 2 error",
			table: ResultTable{
				"platform-1": *platforms.NewResults(err1, err2),
			},
			expected: [][]string{
				header,
				{"platform-1", "error", "err1"},
				{"platform-1", "error", "err2"},
			},
		},
		{
			desc: "1 platform, 1 success, 1 error",
			table: ResultTable{
				"platform-1": *platforms.NewResults(err1, nil),
			},
			expected: [][]string{
				header,
				{"platform-1", "error", "err1"},
				{"platform-1", "ok", ""},
			},
		},
		{
			desc: "2 platform, 2 errors",
			table: ResultTable{
				"platform-1": *platforms.NewResults(err1),
				"platform-2": *platforms.NewResults(err2),
			},
			expected: [][]string{
				header,
				{"platform-1", "error", "err1"},
				{"platform-2", "error", "err2"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			table := c.table.TextTable("ok")

			if !reflect.DeepEqual(table, c.expected) {
				t.Error(cmp.Diff(c.expected, table))
				return
			}
		})
	}
}
