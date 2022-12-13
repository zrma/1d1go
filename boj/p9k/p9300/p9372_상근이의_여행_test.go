package p9300_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9300"
)

func TestSolve9372(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9372")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
3 3
1 2
2 3
1 3
5 4
2 1
2 3
4 3
4 5`,
			`2
4
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p9300.Solve9372(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
