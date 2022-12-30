package p5400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p5k/p5400"
)

func TestSolve5430(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5430")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
RDD
4
[1,2,3,4]
DD
1
[42]
RRD
6
[1,1,2,3,5,8]
D
0
[]
D
1
[1]
D
1
[123]
RDDRDDR
6
[123,123,123,456,456,456]`,
			`[2,1]
error
[1,2,3,5,8]
error
[]
[]
[456,123]
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p5400.Solve5430(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
