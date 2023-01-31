package p10900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10953(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10953")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
1,1
2,3
3,4
9,8
5,2`,
			`2
5
7
17
7
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10953(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
