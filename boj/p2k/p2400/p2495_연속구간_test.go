package p2400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2495(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2495")

	tests := []struct {
		give string
		want string
	}{
		{
			`12345123
17772345
22233331`,
			`1
3
4
`,
		},
		{
			`11111111
12345678
11112222`,
			`8
1
4
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2495(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
