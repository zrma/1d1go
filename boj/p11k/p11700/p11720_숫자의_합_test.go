package p11700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11700"
)

func TestSolve11720(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11720")

	for i, tt := range []struct {
		give string
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
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p11700.Solve11720(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
