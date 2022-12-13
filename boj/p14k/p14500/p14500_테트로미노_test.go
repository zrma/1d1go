package p14500_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p14k/p14500"
)

func TestSolve14500(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14500")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 5
1 2 3 4 5
5 4 3 2 1
2 3 4 5 6
6 5 4 3 2
1 2 1 2 1`,
			"19",
		},
		{
			`4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5`,
			"20",
		},
		{
			`4 10
1 2 1 2 1 2 1 2 1 2
2 1 2 1 2 1 2 1 2 1
1 2 1 2 1 2 1 2 1 2
2 1 2 1 2 1 2 1 2 1`,
			"7",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p14500.Solve14500(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
