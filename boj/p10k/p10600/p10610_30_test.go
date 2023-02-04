package p10600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10610(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10610")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"30", "30"},
		{"102", "210"},
		{"2931", "-1"},
		{"80875542", "88755420"},
		{"1234567890", "9876543210"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10610(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
