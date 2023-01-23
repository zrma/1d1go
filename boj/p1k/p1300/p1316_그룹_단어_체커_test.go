package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1316(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1316")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
happy
new
year`,
			"3",
		},
		{
			`4
aba
abab
abcabc
a`,
			"1",
		},
		{`5
ab
aa
aca
ba
bb`,
			"4",
		},
		{
			`2
yzyzy
zyzyz`,
			"0",
		},
		{
			`1
z`,
			"1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1316(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsGroupWord(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want bool
	}{
		{"happy", true},
		{"new", true},
		{"year", true},
		{"aba", false},
		{"abab", false},
		{"abcabc", false},
		{"a", true},
		{"aa", true},
		{"aca", false},
		{"ba", true},
		{"bb", true},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isGroupWord(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
