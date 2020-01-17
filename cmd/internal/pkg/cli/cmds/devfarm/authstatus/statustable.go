package authstatus

import (
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"sort"
)

func FormatAuthStatusTable(authStatusTable map[platforms.ID]error) [][]string {
	result := make([][]string, len(authStatusTable)+1)
	header := []string{"platform", "auth"}
	result[0] = header
	lineno := 1

	for _, entry := range sortAuthStatusTableEntry(authStatusTable) {
		if entry.authStatus == nil {
			result[lineno] = []string{string(entry.platformID), "success"}
		} else {
			result[lineno] = []string{string(entry.platformID), entry.authStatus.Error()}
		}
		lineno++
	}

	return result
}

type authStatusTableEntry struct {
	platformID platforms.ID
	authStatus error
}

func (entry authStatusTableEntry) Less(another authStatusTableEntry) bool {
	return entry.platformID < another.platformID
}

func sortAuthStatusTableEntry(authStatusTable map[platforms.ID]error) []authStatusTableEntry {
	var entries = make([]authStatusTableEntry, len(authStatusTable))
	i := 0

	for platformID, authStatus := range authStatusTable {
		entries[i] = authStatusTableEntry{platformID, authStatus}
		i++
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Less(entries[j])
	})
	return entries
}
