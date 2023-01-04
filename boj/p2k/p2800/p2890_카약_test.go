package p2800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
)

func TestSolve2890(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2890")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10 10
S.....111F
S....222.F
S...333..F
S..444...F
S.555....F
S666.....F
S.777....F
S..888...F
S...999..F
S........F`,
			`1
2
3
4
5
6
5
4
3
`,
		},
		{
			`10 15
S..........222F
S.....111.....F
S...333.......F
S...555.......F
S.......444...F
S.............F
S......777....F
S..888........F
S........999..F
S...666.......F`,
			`5
1
6
3
6
6
4
7
2
`,
		},
		{
			`10 10
S.....111F
S.....222F
S.....333F
S.....444F
S.....555F
S.....666F
S.....777F
S.....888F
S.....999F
S........F`,
			`1
1
1
1
1
1
1
1
1
`,
		},
		{
			`15 15
S.............F
S.............F
S.............F
S.............F
S.............F
S.....111.....F
S.....222.....F
S.....333.....F
S.....444.....F
S.....555.....F
S.....666.....F
S.....777.....F
S.....888.....F
S.....999.....F
S.............F`,
			`1
1
1
1
1
1
1
1
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2800.Solve2890(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
