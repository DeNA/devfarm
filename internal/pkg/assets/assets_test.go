package assets

import (
	"fmt"
	"testing"
)

func TestAssets(t *testing.T) {
	for _, assetID := range AllAssets {
		t.Run(fmt.Sprintf("Read(%q)", assetID), func(t *testing.T) {
			bin, err := read(assetID)

			if err != nil {
				t.Errorf("got (_, %v), want (_, nil)", err)
				t.Log(AssetNames())
				return
			}

			if len(bin) < 1 {
				t.Errorf("got len(bin) < 1, want len(bin) > 0")
				return
			}
		})
	}
}
