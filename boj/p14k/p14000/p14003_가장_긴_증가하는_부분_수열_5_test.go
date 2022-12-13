package p14000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p14k/p14000"
)

func TestSolve14003(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14003")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
10 20 10 30 20 50`,
			`4
10 20 30 50 `,
		},
		{
			`8
30 1 9 40 7 5 4 90`,
			`4
1 9 40 90 `,
		},
		{
			`5
1 3 5 2 6`,
			`4
1 3 5 6 `,
		},
		{
			`10
1 5 3 4 5 6 7 4 9 10`,
			`8
1 3 4 5 6 7 9 10 `,
		},
		{
			`7
3 1 5 2 3 6 4`,
			`4
1 2 3 4 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p14000.Solve14003(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
