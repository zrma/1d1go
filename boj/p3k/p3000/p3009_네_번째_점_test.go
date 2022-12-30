package p3000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
)

func TestSolve3009(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3009")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 5
5 7
7 5`,
			"7 7",
		},
		{
			`30 20
10 10
10 20`,
			"30 10",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p3000.Solve3009(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
