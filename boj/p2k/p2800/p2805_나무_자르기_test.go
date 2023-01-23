package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2805(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2805")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 7
20 15 10 17`,
			"15",
		},
		{
			`5 20
4 42 40 26 46`,
			"36",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2805(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
