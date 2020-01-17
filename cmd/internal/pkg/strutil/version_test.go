package strutil

import (
	"fmt"
	"testing"
)

func TestLessVersion(t *testing.T) {
	t.Run("a != b", func(t *testing.T) {
		cases := []struct {
			less    string
			greater string
		}{
			{
				less:    "0.0",
				greater: "0.1",
			},
			{
				less:    "0.0",
				greater: "1.0",
			},
			{
				less:    "1",
				greater: "1.1",
			},
			{
				less:    "1",
				greater: "11",
			},
			{
				less:    "ERROR",
				greater: "0",
			},
		}

		for _, c := range cases {
			t.Run(fmt.Sprintf("LessVersion(%s, %s)", c.less, c.greater), func(t *testing.T) {
				if !LessVersion(c.less, c.greater) {
					t.Errorf("got false, want true")
				}
			})

			t.Run(fmt.Sprintf("LessVersion(%s, %s)", c.greater, c.less), func(t *testing.T) {
				if LessVersion(c.greater, c.less) {
					t.Errorf("got true, want false")
				}
			})
		}
	})

	t.Run("a == b", func(t *testing.T) {
		cases := []struct {
			a string
			b string
		}{
			{
				a: "",
				b: "",
			},
			{
				a: "0",
				b: "0",
			},
			{
				a: "0.0",
				b: "0.0",
			},
			{
				a: "1",
				b: "1.0",
			},
			{
				a: "ERROR",
				b: "ALSO_ERROR",
			},
		}

		for _, c := range cases {
			t.Run(fmt.Sprintf("LessVersion(%s, %s)", c.a, c.b), func(t *testing.T) {
				if LessVersion(c.a, c.b) {
					t.Errorf("got false, want true")
				}
			})

			t.Run(fmt.Sprintf("LessVersion(%s, %s)", c.b, c.a), func(t *testing.T) {
				if LessVersion(c.b, c.a) {
					t.Errorf("got true, want false")
				}
			})
		}
	})
}
