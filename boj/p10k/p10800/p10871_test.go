package p10800_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10871(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10871")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`10 5
1 10 4 9 2 3 8 5 7 6`,
			"1 4 2 3 ",
		},
		{
			`10 4
1 10 4 9 2 3 8 5 7 6`,
			"1 2 3 ",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p10800.Solve10871(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
