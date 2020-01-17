package subcmd

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func CommandMissingError(got string, table CommandTable) error {
	message := fmt.Sprintf("No such commands: %s (available: %s)",
		got, strings.Join(availableCommandNames(table), ", "))
	return errors.New(message)
}

func availableCommandNames(table CommandTable) []string {
	keys := make([]string, len(table))
	i := 0

	for key := range table {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	return keys
}
