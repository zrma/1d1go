package p10800_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10828(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10828")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`14
push 1
push 2
top
size
empty
pop
pop
pop
size
empty
pop
push 3
empty
top`,
			`2
2
0
2
1
-1
0
1
-1
0
3
`,
		},
		{
			`7
pop
top
push 123
top
pop
top
pop`,
			`-1
-1
123
123
-1
-1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p10800.Solve10828(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
