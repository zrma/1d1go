package p5600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p5k/p5600"
	"1d1go/utils"
)

func TestSolve5639(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5639")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`50
30
24
5
28
45
98
52
60`,
			`5
28
24
45
30
60
52
98
50
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p5600.Solve5639(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
