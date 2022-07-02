package p9400_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9400"
	"1d1go/utils"
)

func TestSolve9498(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9498")

	for _, tt := range []struct {
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
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p9400.Solve9498(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
