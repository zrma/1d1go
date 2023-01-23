package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1725(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1725")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
2
1
4
5
1
3
3
`,
			"8",
		},
		{"4 1000 1000 1000 1000", "4000"},
		{"12 3 2 3 4 5 4 3 2 3 1 1 1", "18"},
		{"12 28 17 0 7 29 18 16 12 3 30 29 26", "78"},
		{"2 10 7", "14"},
		{"4 8 9 10 1", "24"},
		{"1 1", "1"},
		{"1 100", "100"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1725(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
