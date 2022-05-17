package p1000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1065(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1065")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"110", "99"},
		{"1", "1"},
		{"210", "105"},
		{"1000", "144"},
		{"500", "119"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1000.Solve1065(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
