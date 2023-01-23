package p2900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2920(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2920")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1 2 3 4 5 6 7 8", "ascending"},
		{"8 7 6 5 4 3 2 1", "descending"},
		{"8 1 7 2 6 3 5 4", "mixed"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2920(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
