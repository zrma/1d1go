package p3000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3034(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3034")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 3 4
3
4
5
6
7`,
			`DA
DA
DA
NE
NE
`,
		},
		{
			`2 12 17
21
20`,
			`NE
DA
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3034(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
