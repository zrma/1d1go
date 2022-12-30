package p11000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
)

func TestSolve11053(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11053")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1
1`,
			"1",
		},
		{
			`6
10 20 10 30 20 50`,
			"4",
		},
		{
			`10
1 2 3 4 5 6 7 8 9 10`,
			"10",
		},
		{
			`10
1 5 3 4 5 6 7 4 9 10`,
			"8",
		},
		{
			`10
1 2 3 1 2 3 4 1 2 3`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p11000.Solve11053(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
