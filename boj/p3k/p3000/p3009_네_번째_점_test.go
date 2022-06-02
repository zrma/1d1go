package p3000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
	"1d1go/utils"
)

func TestSolve3009(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3009")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{
			`5 5
5 7
7 5`,
			"7 7",
		},
		{
			`30 20
10 10
10 20`,
			"30 10",
		},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p3000.Solve3009(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
