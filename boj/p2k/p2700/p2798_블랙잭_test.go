package p2700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2798(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2798")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 21
5 6 7 8 9`,
			"21",
		},
		{
			`10 500
93 181 245 214 315 36 185 138 216 295`,
			"497",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2700.Solve2798(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
