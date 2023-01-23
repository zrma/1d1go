package p10800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10818(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10818")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
20 10 35 30 7`,
			"7 35",
		},
		{
			`10
1 2 3 5 4 6 10 9 8 7`,
			"1 10",
		},
		{
			`10
2 3 5 4 6 10 9 8 7 1`,
			"1 10",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10818(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
