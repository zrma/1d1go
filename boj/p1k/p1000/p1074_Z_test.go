package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1074(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1074")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"2 3 1", "11"},
		{"3 1 4", "18"},
		{"3 7 7", "63"},
		{"1 0 0", "0"},
		{"4 7 7", "63"},
		{"10 511 511", "262143"},
		{"10 512 512", "786432"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1074(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
