package strutil

import (
	"strconv"
	"strings"
)

func LessVersion(a string, b string) bool {
	as := strings.Split(string(a), ".")
	bs := strings.Split(string(b), ".")

	al := len(as)
	bl := len(bs)

	var maxLen int
	if al > bl {
		maxLen = al
	} else {
		maxLen = bl
	}

	for i := 0; i < maxLen; i++ {
		var err error
		var ai int
		var bi int

		if i >= al {
			ai = 0
		} else {
			ai, err = strconv.Atoi(as[i])
			if err != nil {
				ai = -1
			}
		}

		if i >= bl {
			bi = 0
		} else {
			bi, err = strconv.Atoi(bs[i])
			if err != nil {
				bi = -1
			}
		}

		if ai == bi {
			continue
		}

		return ai < bi
	}

	return false
}
