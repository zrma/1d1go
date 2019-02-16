package utils

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

type SortAdapter []string

func (a SortAdapter) Len() int           { return len(a) }
func (a SortAdapter) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }
func (a SortAdapter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
