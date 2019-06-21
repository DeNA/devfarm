package formatter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFormatTSV(t *testing.T) {
	cases := []struct {
		table    [][]string
		expected string
	}{
		{
			table:    [][]string{},
			expected: "",
		},
		{
			table: [][]string{
				{},
			},
			expected: "\n",
		},
		{
			table: [][]string{
				{"a"},
			},
			expected: "a\n",
		},
		{
			table: [][]string{
				{"a", "b"},
			},
			expected: "a\tb\n",
		},
		{
			table: [][]string{
				{"a"},
				{"b"},
			},
			expected: "a\nb\n",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("FormatTSV(%q)", c.table), func(t *testing.T) {
			got := FormatTSV(c.table)

			if got != c.expected {
				t.Errorf("got %q, want %q", got, c.expected)
			}
		})
	}
}

func TestPrettyTSV(t *testing.T) {
	cases := []struct {
		table    [][]string
		expected string
	}{
		{
			table:    [][]string{},
			expected: "",
		},
		{
			table: [][]string{
				{},
			},
			expected: "\n",
		},
		{
			table: [][]string{
				{"a"},
			},
			expected: "a\n",
		},
		{
			table: [][]string{
				{"a", "b"},
			},
			expected: "a \tb\n",
		},
		{
			table: [][]string{
				{"a"},
				{"b"},
			},
			expected: "a\nb\n",
		},
		{
			table: [][]string{
				{"a", "cc"},
				{"bb", "d"},
			},
			expected: "a  \tcc\nbb \td \n",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("FormatTSV(%q)", c.table), func(t *testing.T) {
			got := PrettyTSV(c.table)

			if got != c.expected {
				t.Errorf("got %q, want %q", got, c.expected)
			}
		})
	}
}

func TestMaxColumnLen(t *testing.T) {
	cases := []struct {
		table    [][]string
		expected []int
	}{
		{
			table:    [][]string{},
			expected: []int{},
		},
		{
			table:    [][]string{{}},
			expected: []int{},
		},
		{
			table: [][]string{
				{"a", "bb"},
			},
			expected: []int{1, 2},
		},
		{
			table: [][]string{
				{"a"},
				{"bb"},
			},
			expected: []int{2},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("MaxColumnLen(%v)", c.table), func(t *testing.T) {
			got := MaxColumnLen(c.table)

			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}
