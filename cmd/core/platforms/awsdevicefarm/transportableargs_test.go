package awsdevicefarm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestTransportableArgs(t *testing.T) {
	cases := []TransportableArgs{
		{},
		{"a"},
		{"a", "b"},
		{`"`},
		{`'`},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			encoded, encodeErr := EncodeAppArgs(c)
			if encodeErr != nil {
				t.Errorf("got %v, want nil", encodeErr)
				return
			}

			if strings.ContainsAny(encoded, `'$"`) {
				t.Errorf("found unsafe chars in: %q", encoded)
				return
			}

			decoded, decodeErr := DecodeAppArgs(encoded)
			if decodeErr != nil {
				t.Errorf("got %v, want nil", decodeErr)
				return
			}

			if !reflect.DeepEqual(c, decoded) {
				t.Errorf("got %#v, want %#v", c, decoded)
				return
			}
		})
	}
}
