package p9600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9663(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9663")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"2", "0"},
		{"3", "0"},
		{"4", "2"},
		{"5", "10"},
		{"6", "4"},
		{"7", "40"},
		{"8", "92"},
		{"9", "352"},
		{"10", "724"},
		{"11", "2680"},
		{"12", "14200"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve9663(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
