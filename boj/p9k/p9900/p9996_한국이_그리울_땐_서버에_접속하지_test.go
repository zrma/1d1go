package p9900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve996(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9996")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`3
a*d
abcd
anestonestod
facebook`,
			`DA
DA
NE
`,
		},
		{
			`6
h*n
huhovdjestvarnomozedocisvastan
honijezakon
atila
je
bio
hun`,
			`DA
DA
NE
NE
NE
DA
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve9996(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
