package p10800_test

import (
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
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p10800.Solve10844(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
