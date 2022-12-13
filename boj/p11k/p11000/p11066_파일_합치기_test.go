package p11000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
)

func TestSolve11066(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11066")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
4
40 30 30 50
15
1 21 3 4 5 35 5 4 3 5 98 21 14 17 32`,
			`300
864
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p11000.Solve11066(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
