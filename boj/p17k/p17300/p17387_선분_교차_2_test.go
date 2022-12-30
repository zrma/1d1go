package p17300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p17k/p17300"
)

func TestSolve17387(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17387")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 1 5 5
1 5 5 1`,
			"1",
		},
		{
			`1 1 5 5
6 10 10 6`,
			"0",
		},
		{
			`1 1 5 5
5 5 1 1`,
			"1",
		},
		{
			`1 1 5 5
3 3 5 5`,
			"1",
		},
		{
			`1 1 5 5
3 3 1 3`,
			"1",
		},
		{
			`1 1 5 5
5 5 9 9`,
			"1",
		},
		{
			`1 1 5 5
6 6 9 9`,
			"0",
		},
		{
			`1 1 5 5
5 5 1 5`,
			"1",
		},
		{
			`1 1 5 5
6 6 1 5`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p17300.Solve17387(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
