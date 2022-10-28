package p2500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2587(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2587")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10
40
30
60
30`,
			`34
30
`,
		},
		{
			`10
20
35
40
50`,
			`31
35
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2500.Solve2587(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
