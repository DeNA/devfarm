package strutil

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

type OutOfAsciiError struct {
	Input            string
	FirstInvalidChar int32
}

func (e OutOfAsciiError) Error() string {
	return fmt.Sprintf("out of ASCII: %c", e.FirstInvalidChar)
}

// FIXME: It should be IsASCIIPrintable.
func IsASCII(str string) *OutOfAsciiError {
	for _, c := range str {
		if c > unicode.MaxASCII {
			return &OutOfAsciiError{
				Input:            str,
				FirstInvalidChar: c,
			}
		}
	}
	return nil
}

func Indent(str string, width int) string {
	lines := make([]string, 0)
	indent := strings.Repeat(" ", width)

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		lines = append(lines, indent+scanner.Text())
	}

	return strings.Join(lines, "\n")
}
