package p18800_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18800"
)

func TestSolve18870(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18870")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
2 4 -10 4 -9`,
			`2 3 0 3 1 `,
		},
		{
			`6
1000 999 1000 999 1000 999`,
			`1 0 1 0 1 0 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p18800.Solve18870(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
