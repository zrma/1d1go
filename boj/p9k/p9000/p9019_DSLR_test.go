package p9000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9000"
)

func TestSolve9019(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9019")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
1234 3412
1000 1
1 16`,
			`LL
L
DDDD
`,
		},
		{
			`4
1 0
1 10
1 1000
1 1001`,
			`S
L
R
DDSDDSR
`,
		},
		{
			`4
123 1230
123 3012
0 9999
1 1`,
			`L
R
S

`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p9000.Solve9019(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
