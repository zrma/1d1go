package p8300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve8393(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/8393")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"2", "3"},
		{"3", "6"},
		{"10", "55"},
		{"100", "5050"},
		{"1000", "500500"},
		{"10000", "50005000"},
	} {
		t.Run(fmt.Sprintf("loop/%s", tt.give), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve8393(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("arithmetic progression formula/%s", tt.give), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve8393AP(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
