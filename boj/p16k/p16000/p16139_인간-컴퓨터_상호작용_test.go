package p16000_test

import (
	"fmt"
	"strings"
	"testing"

	"1d1go/boj/p16k/p16000"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve16139(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16139")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`seungjaehwang
4
a 0 5
a 0 6
a 6 10
a 7 10`,
			`0
1
2
1
`,
		},
		{
			`asdf
1
a 0 0`,
			`1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := strings.NewReader(tt.give)
			writer := utils.NewStringWriter()

			p16000.Solve16139(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
