package p18800_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18800"
	"1d1go/utils"
)

func TestSolve18870(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18870")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`5
2 4 -10 4 -9`,
			`2 3 0 3 1 `,
		},
		{
			`6
1000 999 1000 999 1000 999`,
			`1 0 1 0 1 0 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p18800.Solve18870(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
