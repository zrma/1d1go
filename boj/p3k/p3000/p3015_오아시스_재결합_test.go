package p3000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
)

func TestSolve3015(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3015")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
2
4
1
2
2
5
1`,
			"10",
		},
		{"3 1 2 1", "2"},
		{"3 2 1 2", "3"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p3000.Solve3015(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
