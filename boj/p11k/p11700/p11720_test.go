package p11700_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11700"
	"1d1go/utils"
)

func TestSolve11720(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11720")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`1
1`,
			"1",
		},
		{
			`5
54321`,
			"15",
		},
		{
			`25
7000000000000000000000000`,
			"7",
		},
		{
			`11
10987654321`,
			"46",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p11700.Solve11720(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
