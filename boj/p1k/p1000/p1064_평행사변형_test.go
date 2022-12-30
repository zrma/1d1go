package p1000_test

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
)

func TestSolve1064(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1064")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`0 0 0 1 1 0`,
			"0.8284271247461898",
		},
		{
			`0 0 4 0 0 3`,
			"4.0",
		},
		{
			`0 0 1 0 47 0`,
			"-1.0",
		},
		{
			`1 2 3 4 8 7`,
			"11.547796284592874",
		},
		{
			`2 -1 -7 2 -1 0`,
			"-1.0",
		},
		{
			`1 1 5 5 6 6`,
			"-1.0",
		},
		{
			`1 1 1 1 5 5`,
			"-1.0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1000.Solve1064(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()

			wantFloat, err := strconv.ParseFloat(tt.want, 64)
			assert.NoError(t, err)

			gotFloat, err := strconv.ParseFloat(got, 64)
			assert.NoError(t, err)

			const epsilon = 1e-9
			assert.InDelta(t, wantFloat, gotFloat, epsilon)
			assert.InEpsilon(t, wantFloat, gotFloat, epsilon)

		})
	}
}
