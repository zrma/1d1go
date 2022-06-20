package p14800_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p14k/p14800"
	"1d1go/utils"
)

func TestSolve14889(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14889")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
0 1 2 3
4 0 5 6
7 1 0 2
3 4 5 0`,
			"0",
		},
		{
			`6
0 1 2 3 4 5
1 0 2 3 4 5
1 2 0 3 4 5
1 2 3 0 4 5
1 2 3 4 0 5
1 2 3 4 5 0`,
			"2",
		},
		{
			`8
0 5 4 5 4 5 4 5
4 0 5 1 2 3 4 5
9 8 0 1 2 3 1 2
9 9 9 0 9 9 9 9
1 1 1 1 0 1 1 1
8 7 6 5 4 0 3 2
9 1 9 1 9 1 0 9
6 5 4 3 2 1 9 0`,
			"1",
		},
		{
			`16
0 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
1 0 1 1 1 1 1 1 1 1 1 1 1 1 1 1
1 1 0 1 1 1 1 1 1 1 1 1 1 1 1 1
1 1 1 0 1 1 1 1 1 1 1 1 1 1 1 1
1 1 1 1 0 1 1 1 1 1 1 1 1 1 1 1
1 1 1 1 1 0 1 1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 0 1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 0 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 0 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1 0 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1 1 0 1 1 1 1 1
1 1 1 1 1 1 1 1 1 1 1 0 1 1 1 1
1 1 1 1 1 1 1 1 1 1 1 1 0 1 1 1
1 1 1 1 1 1 1 1 1 1 1 1 1 0 1 1
1 1 1 1 1 1 1 1 1 1 1 1 1 1 0 1
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 0`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			assert.Eventually(t, func() bool {
				p14800.Solve14889(scanner, writer)

				err := writer.Flush()
				assert.NoError(t, err)

				got := writer.String()
				return assert.Equal(t, tt.want, got)
			}, time.Second, time.Millisecond*100, "시간 초과")
		})
	}
}
