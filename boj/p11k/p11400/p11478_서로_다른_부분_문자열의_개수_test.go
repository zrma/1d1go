package p11400_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11400"
	"1d1go/utils"
)

func TestSolve11478(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11478")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		give string
		want string
	}{
		{"ababc", "12"},
		{"abcde", "15"},
		{"abcdefghij", "55"},
		{"abcdefghab", "52"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11400.Solve11478(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
