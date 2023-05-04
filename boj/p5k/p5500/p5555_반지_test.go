package p5500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve5555(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5555")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`ABCD
3
ABCDXXXXXX
YYYYABCDXX
DCBAZZZZZZ`,
			"2",
		},
		{
			`XYZ
1
ZAAAAAAAXY`,
			"1",
		},
		{
			`PQR
3
PQRAAAAPQR
BBPQRBBBBB
CCCCCCCCCC`,
			"2",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve5555(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
