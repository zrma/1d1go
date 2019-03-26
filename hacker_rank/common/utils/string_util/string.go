package string_util

import (
	"sort"
	"strings"
)

type sortable []byte

func (s sortable) Len() int           { return len(s) }
func (s sortable) Less(i, j int) bool { return s[i] < s[j] }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Sort(s string) string {
	b := sortable(s)
	sort.Sort(b)
	return string(b)
}

type SortAdapter []string

func (a SortAdapter) Len() int           { return len(a) }
func (a SortAdapter) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }
func (a SortAdapter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}