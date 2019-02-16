package utils

import (
	"sort"
	"strings"
)

type sortable []byte

func (s sortable) Len() int           { return len(s) }
func (s sortable) Less(i, j int) bool { return s[i] < s[j] }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func SortString(s string) string {
	b := sortable(s)
	sort.Sort(b)
	return string(b)
}

type SortAdapter []string

func (a SortAdapter) Len() int           { return len(a) }
func (a SortAdapter) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }
func (a SortAdapter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
