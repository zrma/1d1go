package p3000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3029(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3029")

	tests := []struct {
		give string
		want string
	}{
		{
			`20:00:00
04:00:00`,
			"08:00:00",
		},
		{
			`12:34:56
14:36:22`,
			"02:01:26",
		},
		{
			`00:00:00
00:00:01`,
			"00:00:01",
		},
		{
			`23:59:59
00:00:00`,
			"00:00:01",
		},
		{
			`00:00:00
23:59:59`,
			"23:59:59",
		},
		{
			`00:00:01
00:00:00`,
			"23:59:59",
		},
		{
			`00:00:00
00:00:00`,
			"24:00:00",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3029(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
