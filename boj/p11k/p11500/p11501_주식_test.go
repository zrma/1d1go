package p11500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11500"
)

func TestSolve11501(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11501")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
3
10 7 6
3
3 5 9
5
1 1 3 1 2`,
			`0
10
5
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p11500.Solve11501(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
