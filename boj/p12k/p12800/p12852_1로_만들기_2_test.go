package p12800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve12852(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12852")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"2",
			`1
2 1 `,
		},
		{
			"10",
			`3
10 9 3 1 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve12852(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
