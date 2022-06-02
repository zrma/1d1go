package p1400_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1400"
	"1d1go/utils"
)

func TestSolve1427(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1427")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"2143", "4321"},
		{"999998999", "999999998"},
		{"61423", "64321"},
		{"500613009", "965310000"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1400.Solve1427(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
