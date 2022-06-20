package p2500_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2580(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2580")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`0 3 5 4 6 9 2 7 8
7 8 2 1 0 5 6 0 9
0 6 0 2 7 8 1 3 5
3 2 1 0 4 6 8 9 7
8 0 4 9 1 3 5 0 6
5 9 6 8 2 0 4 1 3
9 1 7 6 5 2 0 8 0
6 0 3 7 0 1 9 5 2
2 5 8 3 9 4 7 6 0`,
			`1 3 5 4 6 9 2 7 8
7 8 2 1 3 5 6 4 9
4 6 9 2 7 8 1 3 5
3 2 1 5 4 6 8 9 7
8 7 4 9 1 3 5 2 6
5 9 6 8 2 7 4 1 3
9 1 7 6 5 2 3 8 4
6 4 3 7 8 1 9 5 2
2 5 8 3 9 4 7 6 1
`,
		},
		{
			`0 4 0 0 0 0 0 0 0
0 0 1 0 3 4 6 2 0
6 0 3 0 0 0 0 7 0
0 0 0 4 8 3 5 0 7
0 0 0 0 5 0 0 6 0
0 0 0 0 0 9 0 4 0
0 0 5 0 0 0 0 0 1
8 0 0 5 4 7 3 9 6
0 0 0 0 2 1 0 0 0
`,
			`2 4 8 6 7 5 1 3 9
7 5 1 9 3 4 6 2 8
6 9 3 2 1 8 4 7 5
9 2 6 4 8 3 5 1 7
1 8 4 7 5 2 9 6 3
5 3 7 1 6 9 8 4 2
4 7 5 3 9 6 2 8 1
8 1 2 5 4 7 3 9 6
3 6 9 8 2 1 7 5 4
`,
		},
		{
			`0 0 0 0 0 0 0 0 0
7 8 2 1 3 5 6 4 9
4 6 9 2 7 8 1 3 5
3 2 1 5 4 6 8 9 7
0 0 0 0 0 0 0 0 0
5 9 6 8 2 7 4 1 3
9 1 7 6 5 2 3 8 4
6 4 3 7 8 1 9 5 2
0 0 0 0 0 0 0 0 0`,
			`1 3 5 4 6 9 2 7 8
7 8 2 1 3 5 6 4 9
4 6 9 2 7 8 1 3 5
3 2 1 5 4 6 8 9 7
8 7 4 9 1 3 5 2 6
5 9 6 8 2 7 4 1 3
9 1 7 6 5 2 3 8 4
6 4 3 7 8 1 9 5 2
2 5 8 3 9 4 7 6 1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			assert.Eventually(t, func() bool {
				p2500.Solve2580(scanner, writer)

				err := writer.Flush()
				assert.NoError(t, err)

				got := writer.String()
				return assert.Equal(t, tt.want, got)
			}, time.Second, time.Millisecond*100, "시간 초과")
		})
	}
}
