package p9000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve9012(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
