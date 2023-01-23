package p3700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3733(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3733")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 100
2 7
10 9
10     10
`,
			`50
2
0
0
`,
		},
		{
			`
`,
			"",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3733(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
