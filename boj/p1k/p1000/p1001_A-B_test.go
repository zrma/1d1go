package p1000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1001(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1001")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"3 2", "1"},
		{"1 1", "0"},
		{"9 1", "8"},
		{"1 9", "-8"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1000.Solve1001(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
