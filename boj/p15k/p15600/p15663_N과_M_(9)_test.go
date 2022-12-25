package p15600_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15600"
)

func TestSolve15663(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15663")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 1
4 4 2`,
			`2
4
`,
		},
		{
			`4 2
9 7 9 1`,
			`1 7
1 9
7 1
7 9
9 1
9 7
9 9
`,
		},
		{
			`4 4
1 1 1 1`,
			`1 1 1 1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p15600.Solve15663(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
