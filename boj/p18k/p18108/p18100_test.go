package p18108_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18108"
	"1d1go/utils"
)

func TestSolve18108(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18108")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"2541", "1998"},
		{"543", "0"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p18108.Solve18108(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
