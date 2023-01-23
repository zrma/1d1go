package p3200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3273(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3273")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`9
5 12 7 10 9 1 2 3 11
13`,
			"3",
		},
		{
			`4
1 3 5 1
4`,
			"1",
		},
		{
			`5
1 3 5 3 1
4`,
			"2",
		},
		{
			`5
3 3 5 1 1
4`,
			"2",
		},
	} {
		t.Run(fmt.Sprintf("%d/counting", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3273(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/two point", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3273WithTwoPoints(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
