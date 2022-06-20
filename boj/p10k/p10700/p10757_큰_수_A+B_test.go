package p10700_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10700"
	"1d1go/utils"
)

func TestSolve10757(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10757")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"9223372036854775807 9223372036854775808",
			"18446744073709551615",
		},
		{
			"1 1",
			"2",
		},
		{
			"9 9",
			"18",
		},
		{
			"9 10",
			"19",
		},
		{
			"10 9",
			"19",
		},
		{
			"99999999999999999999999999999999999999999999999 99999999999999999999999999999999999999999999999",
			"199999999999999999999999999999999999999999999998",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p10700.Solve10757(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
