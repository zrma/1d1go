package p1100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
)

func TestSolve1197(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1197")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 3
1 2 1
2 3 2
1 3 3`,
			"3",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1100.Solve1197(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
