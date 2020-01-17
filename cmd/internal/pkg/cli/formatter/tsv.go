package formatter

import (
	"bytes"
	"strings"
)

func FormatTSV(table [][]string) string {
	result := bytes.Buffer{}

	for _, row := range table {
		rowLine := strings.Join(row, "\t")
		result.WriteString(rowLine)
		result.WriteString("\n")
	}

	return result.String()
}

func PrettyTSV(table [][]string) string {
	result := bytes.Buffer{}

	maxLen := MaxColumnLen(table)

	for _, row := range table {
		prettyRow := make([]string, len(row))

		for columnIdx, column := range row {
			prettyRow[columnIdx] = column + strings.Repeat(" ", maxLen[columnIdx]-len(column))
		}

		result.WriteString(strings.Join(prettyRow, " \t"))
		result.WriteString("\n")
	}

	return result.String()
}

func MaxColumnLen(table [][]string) []int {
	if len(table) < 1 {
		return []int{}
	}

	result := make([]int, len(table[0]))

	for columnIdx := range table[0] {
		result[columnIdx] = -1
	}

	for _, row := range table {
		for columnIdx, column := range row {
			columnLen := len(column)
			if result[columnIdx] < columnLen {
				result[columnIdx] = columnLen
			}
		}
	}

	return result
}
