package p24400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve24444(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/24444")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 5 1
1 4
1 2
2 3
2 4
3 4`,
			`1
2
4
3
0
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve24444(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
