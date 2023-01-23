package p9400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9498(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9498")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"100", "A"},
		{"90", "A"},
		{"89", "B"},
		{"80", "B"},
		{"79", "C"},
		{"70", "C"},
		{"69", "D"},
		{"60", "D"},
		{"59", "F"},
		{"0", "F"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve9498(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
