package util

import (
	"sort"
)

// SortByDict
func SortByDict(arr []string) []string {
	sort.Strings(arr)
	return arr
}
