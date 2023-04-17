package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2804(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2804")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`BANANA PIDZAMA`,
			`.P....
.I....
.D....
.Z....
BANANA
.M....
.A....
`,
		},
		{
			`MAMA TATA`,
			`.T..
MAMA
.T..
.A..
`,
		},
		{
			`REPUBLIKA HRVATSKA`,
			`H........
REPUBLIKA
V........
A........
T........
S........
K........
A........
`,
		},
		{
			`AAAA BBBB`,
			`AAAA
B...
B...
B...
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2804(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
