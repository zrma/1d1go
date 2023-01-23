package p1600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1644(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1644")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"20", "0"},
		{"3", "1"},
		{"41", "3"},
		{"53", "2"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1644(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
