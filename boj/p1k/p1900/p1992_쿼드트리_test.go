package p1900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
)

func TestSolve1992(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1992")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`8
11110000
11110000
00011100
00011100
11110000
11110000
11110011
11110011`,
			"((110(0101))(0010)1(0001))",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1900.Solve1992(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
