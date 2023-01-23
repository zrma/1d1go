package p2600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2629(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2629")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
1 4
2
3 2`,
			"Y N",
		},
		{
			`4
2 3 3 3
3
1 4 10`,
			"Y Y N",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2629(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
