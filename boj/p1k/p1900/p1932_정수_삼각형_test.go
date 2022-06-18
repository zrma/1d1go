package p1900_test

import (
	"fmt"
	"testing"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve1932(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1932")

	for i, tt := range []struct {
		s    string
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
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1900.Solve1932(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
