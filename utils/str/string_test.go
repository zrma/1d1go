package str

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Parallel()

	t.Log("Sort function sorts a string well")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		s    string
		want string
	}{
		{"dcba", "abcd"},
		{"ffbbaa", "aabbff"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Sort(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSortAdapter(t *testing.T) {
	t.Parallel()

	t.Log("SortAdapter structure sorts string slices well")

	for i, tt := range []struct {
		strArr []string
		want   []string
	}{
		{
			[]string{"a", "d", "b", "c"},
			[]string{"a", "b", "c", "d"},
		},
		{
			[]string{"a", "d", "bc", "ba"},
			[]string{"a", "ba", "bc", "d"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := make([]string, len(tt.strArr))
			copy(got, tt.strArr)
			sort.Sort(SortAdapter(got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()

	t.Run("Reverse function reverses given string", func(t *testing.T) {
		const (
			s    = "abc"
			want = "cba"
		)
		got := Reverse(s)
		assert.Equal(t, want, got)
	})

	t.Run("no error with empty string", func(t *testing.T) {
		assert.NotPanics(t, func() {
			const emptyStr = ""
			got := Reverse(emptyStr)
			assert.Equal(t, emptyStr, got)
		})
	})
}
