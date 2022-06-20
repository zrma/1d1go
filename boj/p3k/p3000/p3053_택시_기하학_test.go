package p3000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
	"1d1go/utils"
)

func TestSolve3053(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3053")

	for _, tt := range []struct {
		give string
		want string
	}{
		{
			"1",
			`3.141593
2.000000
`,
		},
		{
			"21",
			`1385.442360
882.000000
`,
		},
		{
			"42",
			`5541.769441
3528.000000
`,
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p3000.Solve3053(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
