package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1302(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1302")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`5
top
top
top
top
kimtop`,
			`top`,
		},
		{
			`9
table
chair
table
table
lamp
door
lamp
table
chair`,
			`table`,
		},
		{
			`6
a
a
a
b
b
b`,
			`a`,
		},
		{
			`8
icecream
peanuts
peanuts
chocolate
candy
chocolate
icecream
apple`,
			`chocolate`,
		},
		{
			`1
soul`,
			`soul`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1302(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
