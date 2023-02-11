package p1400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1439(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1439")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"0001100", "1"},
		{"11111", "0"},
		{"00000001", "1"},
		{"11001100110011000001", "4"},
		{"11101101", "2"},
		{"0", "0"},
		{"1", "0"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1439(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
