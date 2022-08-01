package p2100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2100"
	"1d1go/utils"
)

func TestSolve2110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2110")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 3
1
2
8
4
9`,
			"3",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2100.Solve2110(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
