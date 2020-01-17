package planfile

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/strutil"
	"strings"
)

type SemanticError struct {
	errors       []error
	locationHint string
}

func (e SemanticError) Error() string {
	s := make([]string, len(e.errors))

	for i, err := range e.errors {
		s[i] = err.Error()
	}

	return fmt.Sprintf(
		"Planfile error at %s\n%s",
		e.locationHint,
		strutil.Indent(strings.Join(s, "\n"), 4),
	)
}
