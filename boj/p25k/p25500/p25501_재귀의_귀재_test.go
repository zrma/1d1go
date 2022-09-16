package p25500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p25k/p25500"
	"1d1go/utils"
)

func TestSolve25501(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/25501")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
AAA
ABBA
ABABA
ABCA
PALINDROME`,
			`1 2
1 3
1 3
0 2
0 1
`,
		},
		{
			`1
A`,
			`1 1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p25500.Solve25501(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
