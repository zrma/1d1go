package p11500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11586(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11586")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`8
OOOOOOOO
OKKOOEEO
OKKOOEEO
OOOSSOOO
OOOSSOOO
OAOOOOAO
OOAAAAOO
OOOOOOOO
3`,
			`OOOOOOOO
OOAAAAOO
OAOOOOAO
OOOSSOOO
OOOSSOOO
OKKOOEEO
OKKOOEEO
OOOOOOOO
`,
		},
		{
			`8
OOOOOOOO
OKKOOEEO
OKKOOEEO
OOOSSOOO
OOOSSOOO
OAOOOOAO
OOAAAAOO
OOOOOOOO
1`,
			`OOOOOOOO
OKKOOEEO
OKKOOEEO
OOOSSOOO
OOOSSOOO
OAOOOOAO
OOAAAAOO
OOOOOOOO
`,
		},
		{
			`8
OOOOOOOO
OKKOOEEO
OKKOOEEO
OOOSSOOO
OOOSSOOO
OAOOOOAO
OOAAAAOO
OOOOOOOO
2`,
			`OOOOOOOO
OEEOOKKO
OEEOOKKO
OOOSSOOO
OOOSSOOO
OAOOOOAO
OOAAAAOO
OOOOOOOO
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11586(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
