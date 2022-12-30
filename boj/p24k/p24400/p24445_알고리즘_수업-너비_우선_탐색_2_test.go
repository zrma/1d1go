package p24400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p24k/p24400"
)

func TestSolve24445(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/24445")

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
3
4
2
0
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p24400.Solve24445(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
