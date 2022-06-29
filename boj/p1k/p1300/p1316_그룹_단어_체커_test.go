package p1300_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1300"
	"1d1go/utils"
)

func TestSolve1316(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1316")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
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
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1316(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
