package p11700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11727(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11727")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"2", "3"},
		{"8", "171"},
		{"12", "2731"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11727(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
