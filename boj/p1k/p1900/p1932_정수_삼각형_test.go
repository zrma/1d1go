package p1900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
)

func TestSolve1932(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1932")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
7
3 8
8 1 0
2 7 4 4
4 5 2 6 5`,
			"30",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1900.Solve1932(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
