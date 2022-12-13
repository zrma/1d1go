package p1900_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
)

func TestSolve1949(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1949")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
1000 3000 4000 1000 2000 2000 7000
1 2
2 3
4 3
4 5
6 2
6 7`,
			"14000",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1900.Solve1949(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
