package p18100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18100"
)

func TestSolve18111(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18111")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 4 99
0 0 0 0
0 0 0 0
0 0 0 1`,
			"2 0",
		},
		{
			`3 4 1
64 64 64 64
64 64 64 64
64 64 64 63`,
			"1 64",
		},
		{
			`3 4 0
64 64 64 64
64 64 64 64
64 64 64 63`,
			"22 63",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p18100.Solve18111(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
