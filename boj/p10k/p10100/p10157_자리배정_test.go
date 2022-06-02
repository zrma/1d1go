package p10100_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10100"
	"1d1go/utils"
)

func TestSolve10157(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10157")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`7 6
11`,
			"6 6",
		},
		{
			`7 6
87`,
			"0",
		},
		{
			`100 100
3000`,
			"9 64",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p10100.Solve10157(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
