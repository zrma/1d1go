package p2900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2935(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2935")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1000
*
100`,
			`100000`,
		},
		{
			`10000
+
10`,
			`10010`,
		},
		{
			`10
+
1000`,
			`1010`,
		},
		{
			`1
*
1000`,
			`1000`,
		},
		{
			`1000
+
1000`,
			`2000`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2935(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
