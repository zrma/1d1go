package p2000_test

import (
	"1d1go/boj/p2k/p2000"
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSolve2012(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2012")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
1
5
3
1
2`,
			"3",
		},
		{
			`5
1
3
3
3
3`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2000.Solve2012(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
