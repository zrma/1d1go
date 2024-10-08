package p2400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2470(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2470")

	tests := []struct {
		give string
		want string
	}{
		{
			`5
-2 4 -99 -1 98`,
			"-99 98",
		},
		{
			`5
-2 3 -99 1 97`,
			"-2 3",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2470(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
