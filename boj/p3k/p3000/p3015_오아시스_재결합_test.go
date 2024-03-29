package p3000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3015(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3015")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
2
4
1
2
2
5
1`,
			"10",
		},
		{"3 1 2 1", "2"},
		{"3 2 1 2", "3"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3015(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
