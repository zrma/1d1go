package p11000_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
)

func TestSolve11050(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11050")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1 0", "1"},
		{"1 1", "1"},
		{"1 2", "0"},
		{"5 2", "10"},
		{"6 1", "6"},
		{"6 2", "15"},
		{"6 3", "20"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11000.Solve11050(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
