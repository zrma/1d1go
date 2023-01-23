package p16000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve16139(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16139")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`seungjaehwang
4
a 0 5
a 0 6
a 6 10
a 7 10`,
			`0
1
2
1
`,
		},
		{
			`asdf
1
a 0 0`,
			`1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve16139(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
