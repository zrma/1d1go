package p2700

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2711(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2711")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`4
4 MISSPELL
1 PROGRAMMING
7 CONTEST
3 BALLOON`,
			`MISPELL
ROGRAMMING
CONTES
BALOON
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(bytes.NewBufferString(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2711(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
