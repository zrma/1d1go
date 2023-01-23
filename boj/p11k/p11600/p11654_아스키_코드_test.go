package p11600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11654(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11654")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"A", "65"},
		{"B", "66"},
		{"C", "67"},
		{"0", "48"},
		{"1", "49"},
		{"9", "57"},
		{"a", "97"},
		{"z", "122"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11654(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
