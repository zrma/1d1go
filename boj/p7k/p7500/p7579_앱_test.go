package p7500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve7579(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7579")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 60
30 10 20 35 40
3 0 3 5 4`,
			"6",
		},
		{
			`5 60
30 10 20 35 40
3 0 3 5 2`,
			"5",
		},
		{
			`5 60
10 10 20 10 10
1000 1000 1000 1000 6000`,
			"10000",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve7579(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
