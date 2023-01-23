package p1800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1874(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1874")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`8
4
3
6
8
7
5
2
1`,
			`+
+
+
+
-
-
+
+
-
+
+
-
-
-
-
-
`,
		},
		{
			`5
1
2
5
3
4`,
			"NO",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1874(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
