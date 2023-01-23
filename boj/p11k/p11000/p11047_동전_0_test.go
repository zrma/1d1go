package p11000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11047(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11047")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10 4200
1
5
10
50
100
500
1000
5000
10000
50000`,
			"6",
		},
		{
			`10 4790
1
5
10
50
100
500
1000
5000
10000
50000`,
			"12",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11047(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
