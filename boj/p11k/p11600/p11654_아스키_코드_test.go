package p11600_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11600"
	"1d1go/utils"
)

func TestSolve11654(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11654")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"A", "65"},
		{"B", "66"},
		{"C", "67"},
		{"0", "48"},
		{"1", "49"},
		{"9", "57"},
		{"a", "97"},
		{"z", "122"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p11600.Solve11654(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
