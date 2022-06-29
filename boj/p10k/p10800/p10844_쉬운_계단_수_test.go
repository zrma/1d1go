package p10800_test

import (
	"bufio"
	"strings"
	"testing"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve10844(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10844")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1", "9"},
		{"2", "17"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p10800.Solve10844(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
