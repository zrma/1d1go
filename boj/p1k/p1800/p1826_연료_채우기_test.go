package p1800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1800"
)

func TestSolve1826(t *testing.T) {
	t.Log("https://github.com/golang/go/issues/1826")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
4 4
5 2
11 5
15 10
25 10`,
			"3",
		},
		{
			`4
4 4
5 2
11 5
15 10
25 5`,
			"4",
		},
		{
			`4
4 4
5 2
11 5
15 10
25 4`,
			"-1",
		},
		{
			`4
4 20
5 2
11 5
15 10
25 5`,
			"1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1800.Solve1826(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
