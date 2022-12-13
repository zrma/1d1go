package p2900_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2900"
)

func TestSolve2908(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2908")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"734 893", "437"},
		{"221 231", "132"},
		{"839 237", "938"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2900.Solve2908(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
