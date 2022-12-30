package p25300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p25k/p25300"
)

func TestSolve25304(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/25304")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`260000
4
20000 5
30000 2
10000 6
5000 8`,
			"Yes",
		},
		{
			`250000
4
20000 5
30000 2
10000 6
5000 8`,
			"No",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p25300.Solve25304(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
