package p3000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
)

func TestSolve3036(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3036")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
8 4 2`,
			`2/1
4/1
`,
		},
		{
			`4
12 3 8 4`,
			`4/1
3/2
3/1
`,
		},
		{
			`4
300 1 1 300`,
			`300/1
300/1
1/1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p3000.Solve3036(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
