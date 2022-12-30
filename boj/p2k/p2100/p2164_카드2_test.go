package p2100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2100"
)

func TestSolve2164(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2164")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"4", "4"},
		{"6", "4"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2100.Solve2164(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
