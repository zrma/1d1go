package p15600_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15600"
)

func TestSolve15681(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15681")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`9 5 3
1 3
4 3
5 4
5 6
6 7
2 3
9 6
6 8
5
4
8`,
			`9
4
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p15600.Solve15681(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
