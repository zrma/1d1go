package p12000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve12018(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12018")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 76
5 4
36 25 1 36 36
4 4
30 24 25 20
6 4
36 36 36 36 36 36
2 4
3 7
5 4
27 15 26 8 14`,
			"4",
		},
		{
			`5 999999
5 4
36 25 1 36 36
4 4
30 24 25 20
6 4
36 36 36 36 36 36
2 4
3 7
5 4
27 15 26 8 14`,
			"5",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve12018(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
