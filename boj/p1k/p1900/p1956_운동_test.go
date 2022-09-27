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

func TestSolve1956(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1956")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 4
1 2 1
3 2 1
1 3 5
2 3 2`,
			"3",
		},
		{
			`3 0`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1900.Solve1956(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
