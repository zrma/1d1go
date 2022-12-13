package p1100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
)

func TestSolve1110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1110")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"26", "4"},
		{"55", "3"},
		{"1", "60"},
		{"0", "1"},
		{"71", "12"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1100.Solve1110(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
