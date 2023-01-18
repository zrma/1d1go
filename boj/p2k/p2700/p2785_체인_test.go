package p2700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
)

func TestSolve2785(t *testing.T) {
	t.Log("https://en.wikipedia.org/wiki/P2700")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
3 3`,
			"1",
		},
		{
			`3
1 1 1`,
			"1",
		},
		{
			`5
4 3 5 7 9`,
			"3",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2700.Solve2785(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
