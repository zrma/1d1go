package p14600_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p14k/p14600"
	"1d1go/utils"
)

func TestSolve14681(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14681")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`12
5`,
			"1",
		},
		{
			`-12
5`,
			"2",
		},
		{
			`-9
-13`,
			"3",
		},
		{
			`9
-13`,
			"4",
		},
		{
			`0
0`,
			"-1",
		},
		{
			`0
-1`,
			"-1",
		},
		{
			`369
0`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p14600.Solve14681(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
