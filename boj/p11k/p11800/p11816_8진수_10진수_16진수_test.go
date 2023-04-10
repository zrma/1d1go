package p11800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11816(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11816")

	tests := []struct {
		give string
		want string
	}{
		{"10", "10"},
		{"010", "8"},
		{"0x10", "16"},
		{"0x3f6", "1014"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11816(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
