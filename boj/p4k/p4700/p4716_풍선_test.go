package p4700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve4716(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4716")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 15 35
10 20 10
10 10 30
10 40 10
3 15 35
10 20 10
10 10 30
10 40 10
0 0 0`,
			`300
300
`,
		},
		{
			`3 10 20
10 0 10
10 0 10
10 0 10
3 20 10
10 10 0
10 10 0
10 10 0
0 0 0`,
			`200
200
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve4716(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			git := buf.String()
			assert.Equal(t, tt.want, git)
		})
	}
}
