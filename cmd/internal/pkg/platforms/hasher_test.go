package platforms

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewCRC32Hasher(t *testing.T) {
	cases := []struct {
		input       []byte
		expected    uint32
		expectedErr bool
	}{
		{
			input:       []byte{},
			expected:    0x0,
			expectedErr: false,
		},
		{
			input:       []byte{},
			expected:    0x0,
			expectedErr: false,
		},
		{
			input:       []byte("hello"),
			expected:    0x3610a686,
			expectedErr: false,
		},
	}

	for _, c := range cases {
		hash := NewCRC32Hasher()

		t.Run(fmt.Sprintf("hash := newCRC32Hasher(); hash(%v)", c.input), func(t *testing.T) {
			buf := &bytes.Buffer{}
			buf.Write(c.input)

			got, err := hash(buf)
			if c.expectedErr {
				if err == nil {
					t.Errorf("got (_, nil), want (_, error)")
				}
			} else {
				if err != nil {
					t.Errorf("got (_, %v), want (_, nil)", err)
				} else if got != c.expected {
					t.Errorf("got (%v, nil), want (%v, nil)", got, c.expected)
				}
			}
		})
	}
}
