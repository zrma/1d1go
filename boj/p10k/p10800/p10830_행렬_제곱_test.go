package p10800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10830(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10830")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2 1
1 2
3 4`,
			`1 2
3 4
`,
		},
		{
			`2 5
1 2
3 4`,
			`69 558
337 406
`,
		},
		{
			`3 3
1 2 3
4 5 6
7 8 9`,
			`468 576 684
62 305 548
656 34 412
`,
		},
		{
			`5 10
1 0 0 0 1
1 0 0 0 1
1 0 0 0 1
1 0 0 0 1
1 0 0 0 1`,
			`512 0 0 0 512
512 0 0 0 512
512 0 0 0 512
512 0 0 0 512
512 0 0 0 512
`},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p10800.Solve10830(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
