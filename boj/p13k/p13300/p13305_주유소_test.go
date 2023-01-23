package p13300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve13305(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/13305")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
2 3 1
5 2 4 1`,
			"18",
		},
		{
			`4
3 3 4
1 1 1 1`,
			"10",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve13305(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
