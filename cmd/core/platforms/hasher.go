package platforms

import (
	"hash/crc32"
	"io"
)

type Hasher32 func(reader io.Reader) (uint32, error)

// NOTE: Should use more longer hash values...? At least one pair exists if your app size is 1GB.
func NewCRC32Hasher() Hasher32 {
	return func(reader io.Reader) (uint32, error) {
		h := crc32.NewIEEE()

		if _, err := io.Copy(h, reader); err != nil {
			return 0, err
		}

		return h.Sum32(), nil
	}
}
