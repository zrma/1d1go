package p11700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11724(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11724")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6 5
1 2
2 5
5 1
3 4
4 6`,
			"2",
		},
		{
			`6 8
1 2
2 5
5 1
3 4
4 6
5 4
2 4
2 3`,
			"1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11724(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
