package p11900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11945(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11945")

	tests := []struct {
		give string
		want string
	}{
		{
			`5 7
0010000
0111010
1111111
0111010
0010000`,
			`0000100
0101110
1111111
0101110
0000100
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11945(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
