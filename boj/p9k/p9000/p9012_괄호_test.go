package p9000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9000"
	"1d1go/utils"
)

func TestSolve9012(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9012")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
(())())
(((()())()
(()())((()))
((()()(()))(((())))()
()()()()(()()())()
(()((())()(`,
			`NO
NO
YES
NO
YES
NO
`,
		},
		{
			`3
((
))
())(()`,
			`NO
NO
NO
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p9000.Solve9012(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
