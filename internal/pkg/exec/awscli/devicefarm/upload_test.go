package devicefarm

import (
	"encoding/json"
	"testing"
)

func TestUpload_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		json     []byte
		expected Upload
	}{
		{
			json:     []byte(initializedUploadJSONExample),
			expected: initializedUploadExample,
		},
		{
			json:     []byte(succeededUploadJSONExample),
			expected: succeededUploadExample,
		},
	}

	for _, c := range cases {
		t.Run("json.Unmarshal(%q, _)", func(t *testing.T) {
			var got Upload

			if err := json.Unmarshal([]byte(c.json), &got); err != nil {
				t.Errorf("got %v, want nil", err)
				return
			}

			if got != c.expected {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}
