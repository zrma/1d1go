package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1086(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1086")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
3
2
1
2`,
			"1/3",
		},
		{
			`5
10
100
1000
10000
100000
10`,
			"1/1",
		},
		{
			`5
11
101
1001
10001
100001
10`,
			"0/1",
		},
		{
			`9
13
10129414190271203
102
102666818896
1216
1217
1218
101278001
1000021412678412681
21`,
			"5183/36288",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1086(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
