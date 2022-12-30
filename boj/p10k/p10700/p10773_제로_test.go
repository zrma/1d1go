package p10700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10700"
)

func TestSolve10773(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10773")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
3
0
4
0`,
			"0",
		},
		{
			`10
1
3
5
4
0
0
7
0
0
6`,
			"7",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p10700.Solve10773(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
