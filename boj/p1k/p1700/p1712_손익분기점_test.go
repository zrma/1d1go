package p1700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1700"
	"1d1go/utils"
)

func TestSolve1712(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1712")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1000 70 170", "11"},
		{"3 2 1", "-1"},
		{"2100000000 9 10", "2100000001"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1700.Solve1712(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
