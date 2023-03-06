package p2600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2675(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2675")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`3
3 ABC
5 /HTP
2 \$%*+-./`,
			`AAABBBCCC
/////HHHHHTTTTTPPPPP
\\$$%%**++--..//
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2675(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
