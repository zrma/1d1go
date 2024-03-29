package p5600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve5639(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5639")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`50
30
24
5
28
45
98
52
60`,
			`5
28
24
45
30
60
52
98
50
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve5639(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
