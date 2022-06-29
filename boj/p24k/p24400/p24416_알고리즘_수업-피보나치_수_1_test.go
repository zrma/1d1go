package p24400_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p24k/p24400"
	"1d1go/utils"
)

func TestSolve24416(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/24416")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"5", "5 3"},
		{"30", "832040 28"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p24400.Solve24416(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
