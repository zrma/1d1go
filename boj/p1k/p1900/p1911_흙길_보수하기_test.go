package p1900_test

import (
	"1d1go/boj/p1k/p1900"
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSolve1911(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1911")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 3
1 6
13 17
8 12`,
			"5",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1900.Solve1911(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
