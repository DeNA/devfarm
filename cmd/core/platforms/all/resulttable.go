package all

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/strutil"
	"sort"
	"strings"
	"sync"
)

type ResultTable map[platforms.ID]platforms.Results

func (t ResultTable) Err() error {
	messages := make([]string, 0)

	for _, results := range t {
		for _, err := range results.ErrorsOnlyNotNil() {
			messages = append(messages, err.Error())
		}
	}

	count := len(messages)
	if count > 0 {
		return fmt.Errorf(
			"%d errors ocurred:\n%s",
			count,
			strutil.Indent(strings.Join(messages, "\n"), 4),
		)
	}
	return nil
}

func (t ResultTable) TextTable(successMsg string) [][]string {
	body := make([][]string, 0)

	for platformID, results := range t {
		for _, err := range results {
			var row []string
			if err != nil {
				row = []string{string(platformID), "error", err.Error()}
			} else {
				row = []string{string(platformID), successMsg, ""}
			}
			body = append(body, row)
		}
	}
	sort.Slice(body, func(i, j int) bool {
		if body[i][0] != body[j][0] {
			return body[i][0] < body[j][0]
		}
		if body[i][1] != body[j][1] {
			return body[i][1] < body[j][1]
		}
		if body[i][2] != body[j][2] {
			return body[i][2] < body[j][2]
		}
		return false
	})

	result := make([][]string, 1+len(body))
	result[0] = []string{"platform", "status", "note"}
	for i, row := range body {
		result[i+1] = row
	}
	return result
}

type ResultTableBuilder struct {
	mu       sync.Mutex
	mutTable map[platforms.ID]*platforms.Results
}

func NewResultTableBuilder() *ResultTableBuilder {
	return &ResultTableBuilder{
		mutTable: make(map[platforms.ID]*platforms.Results),
	}
}

func (t *ResultTableBuilder) Build() ResultTable {
	table := make(ResultTable, len(t.mutTable))
	for platformID, results := range t.mutTable {
		table[platformID] = *results
	}
	return table
}

func (t *ResultTableBuilder) AddErrors(platformID platforms.ID, errors ...error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	results, ok := t.mutTable[platformID]
	if !ok {
		results = platforms.NewResults()
		t.mutTable[platformID] = results
	}
	results.AddErrorOrNils(errors...)
}

func (t *ResultTableBuilder) AddSuccesses(platformID platforms.ID, delta int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	results, ok := t.mutTable[platformID]
	if !ok {
		results = platforms.NewResults()
		t.mutTable[platformID] = results
	}
	results.AddSuccesses(delta)
}
