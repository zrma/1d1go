package p1500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1520(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1520")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 5
50 45 37 32 30
35 50 40 20 25
30 30 25 17 28
27 24 22 15 10`,
			"3",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1520(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
