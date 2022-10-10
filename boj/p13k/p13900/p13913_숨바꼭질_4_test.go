package p13900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p13k/p13900"
	"1d1go/utils"
)

func TestSolve13913(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/13913")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"5 17",
			`4
5 4 8 16 17 `,
		},
		{
			"15 0",
			`15
15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0 `,
		},
		{
			"0 15",
			`6
0 1 2 4 8 16 15 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p13900.Solve13913(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
