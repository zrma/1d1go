package p17200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve17299(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17299")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
1 1 2 3 4 2 1`,
			`-1 -1 1 2 2 1 -1 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve17299(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
