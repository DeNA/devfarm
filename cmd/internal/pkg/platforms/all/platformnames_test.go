package all

import (
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestPlatformNames(t *testing.T) {
	ps := NewPlatforms(platforms.AnyBag())
	expected := make(map[platforms.ID]bool, len(ps.table))
	for platformID := range ps.table {
		expected[platformID] = true
	}

	got := ValidPlatformIDs

	if !reflect.DeepEqual(got, expected) {
		t.Error(cmp.Diff(expected, got))
		return
	}
}
