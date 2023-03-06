package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1373(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1373")

	tests := []struct {
		give string
		want string
	}{
		{"11001100", "314"},
		{"11011", "33"},
		{"11111111", "377"},
		{"11", "3"},
		{"1", "1"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1373(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
