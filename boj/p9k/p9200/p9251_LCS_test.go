package p9200_test

import (
	"fmt"
	"testing"

	"1d1go/boj/p9k/p9200"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve9251(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9251")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`ACAYKP
CAPCAK`,
			"4",
		},
		{
			`ABCDEF
GBCDEF`,
			"5",
		},
		{
			`ABCDEF
GBCDFE`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p9200.Solve9251(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
