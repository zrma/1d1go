package p1000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1003(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1003")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
0
1
3`,
			`1 0
0 1
1 2
`,
		},
		{
			`2
6
22`,
			`5 8
10946 17711
`,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1000.Solve1003(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
