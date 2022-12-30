package p20100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p20k/p20100"
)

func TestSolve20149(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/20149")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 1 5 5
1 5 5 1`,
			`1
3 3`,
		},
		{
			`1 1 5 5
6 10 10 6`,
			`0
`,
		},
		{
			`1 1 5 5
5 5 1 1`,
			`1
`,
		},
		{
			`1 1 5 5
3 3 5 5`,
			`1
`,
		},
		{
			`1 1 5 5
3 3 1 3`,
			`1
3 3`,
		},
		{
			`1 1 5 5
5 5 9 9`,
			`1
5 5`,
		},
		{
			`1 1 5 5
6 6 9 9`,
			`0
`,
		},
		{
			`1 1 5 5
5 5 1 5`,
			`1
5 5`,
		},
		{
			`1 1 5 5
6 6 1 5`,
			`0
`,
		},
		{
			`0 0 10 0
5 0 15 0`,
			`1
`,
		},
		{
			`-4 5 -4 -4
5 -5 -5 5`,
			`1
-4 4`,
		},
		{
			`5 5 1 1
5 5 9 9`,
			`1
5 5`,
		},
		{
			`0 0 1 1
0 0 1 1`,
			`1
`,
		},
		{
			`0 0 1 1
1 1 0 0`,
			`1
`,
		},
		{
			`0 0 1 1
2 2 1 1`,
			`1
1 1`,
		},
		{
			`0 0 0 1
0 2 0 1`,
			`1
0 1`,
		},
		{
			`0 0 0 1
0 0 0 -1`,
			`1
0 0`,
		},
		{
			`0 0 0 1
0 2 0 3`,
			`0
`,
		},
		{
			`-1 -1 1 1
1 -1 -1 1`,
			`1
0 0`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p20100.Solve20149(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}

	t.Run("float64", func(t *testing.T) {
		const (
			give = `2 8 9 23
1 10 9 8`
			wantVal = 1
			wantX   = 2.7313432835820897
			wantY   = 9.5671641791044770
		)

		reader := bufio.NewReader(strings.NewReader(give))
		buf := new(strings.Builder)
		writer := bufio.NewWriter(buf)

		p20100.Solve20149(reader, writer)

		err := writer.Flush()
		assert.NoError(t, err)

		got := buf.String()
		stringReader := strings.NewReader(got)
		var gotVal int
		var gotX, gotY float64

		_, err = fmt.Fscan(stringReader, &gotVal, &gotX, &gotY)
		assert.NoError(t, err)

		assert.Equal(t, wantVal, gotVal)

		const epsilon = 1e-9
		assert.InDelta(t, wantX, gotX, epsilon)
		assert.InDelta(t, wantY, gotY, epsilon)
	})
}
