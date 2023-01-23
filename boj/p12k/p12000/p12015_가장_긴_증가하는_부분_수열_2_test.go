package p12000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve12015(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12015")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
10 20 10 30 20 50`,
			"4",
		},
		{
			`8
30 1 9 40 7 5 4 90`,
			"4",
		},
		{
			`5
1 3 5 2 6`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve12015(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
