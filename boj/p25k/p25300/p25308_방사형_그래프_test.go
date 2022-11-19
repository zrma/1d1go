package p25300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p25k/p25300"
	"1d1go/utils"
)

func TestSolve25308(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/25308")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 1 1 1 1 1 1 1`,
			"40320",
		},
		{
			`6 7 7 8 9 10 11 13`,
			"7712",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p25300.Solve25308(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
