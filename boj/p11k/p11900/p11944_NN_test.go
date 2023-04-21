package p11900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11944(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11944")

	tests := []struct {
		give string
		want string
	}{
		{"20 16", "2020202020202020"},
		{"20 20", "20202020202020202020"},
		{"20 30", "202020202020202020202020202020"},
		{"1 100", "1"},
		{"100 1", "1"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11944(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
