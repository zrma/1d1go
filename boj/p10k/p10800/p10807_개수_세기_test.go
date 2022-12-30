package p10800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10807(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10807")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`11
1 4 1 2 4 2 4 2 3 4 4
2`,
			"3",
		},
		{
			`11
1 4 1 2 4 2 4 2 3 4 4
5`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p10800.Solve10807(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
