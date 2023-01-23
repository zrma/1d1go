package p2100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2166(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2166")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
0 0
0 10
10 10
10 0`,
			"100.0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2166(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
