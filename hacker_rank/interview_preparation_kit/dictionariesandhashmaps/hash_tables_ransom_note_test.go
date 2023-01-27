package dictionariesandhashmaps

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMagazine(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/ctci-ransom-note/problem")

	for i, tt := range []struct {
		magazine, note []string
		want           string
	}{
		{
			magazine: []string{"give", "me", "one", "grand", "today", "night"},
			note:     []string{"give", "one", "grand", "today"},
			want:     "Yes",
		},
		{
			magazine: []string{"two", "times", "three", "is", "not", "four"},
			note:     []string{"two", "times", "two", "is", "four"},
			want:     "No",
		},
		{
			magazine: []string{"ive", "got", "a", "lovely", "brunch", "of", "coconuts"},
			note:     []string{"ive", "got", "some", "coconuts"},
			want:     "No",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			fmtPrint = func(a ...any) (n int, err error) {
				return fmt.Fprint(writer, a...)
			}
			defer func() { fmtPrint = fmt.Print }()

			checkMagazine(tt.magazine, tt.note)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
