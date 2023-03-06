package p2000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2012(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2012")

	tests := []struct {
		give string
		want string
	}{
		{
			`5
1
5
3
1
2`,
			"3",
		},
		{
			`5
1
3
3
3
3`,
			"4",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2012(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
