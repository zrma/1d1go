package p3000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3053(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3053")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"1",
			`3.141593
2.000000
`,
		},
		{
			"21",
			`1385.442360
882.000000
`,
		},
		{
			"42",
			`5541.769441
3528.000000
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3053(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
