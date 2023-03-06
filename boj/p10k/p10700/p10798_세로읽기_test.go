package p10700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10798(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10798")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`ABCDE
abcde
01234
FGHIJ
fghij`,
			"Aa0FfBb1GgCc2HhDd3IiEe4Jj",
		},
		{
			`AABCDD
afzz
09121
a8EWg6
P5h3kx`,
			"Aa0aPAf985Bz1EhCz2W3D1gkD6x",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10798(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
