package p2500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2533(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2533")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`8
1 2
1 3
1 4
2 5
2 6
4 7
4 8`,
			"3",
		},
		{
			`9
1 2
1 3
2 4
3 5
3 6
4 7
4 8
4 9`,
			"3",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2533(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
