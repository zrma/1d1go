package p16900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve16953(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16953")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"2 3", "-1"},
		{"2 162", "5"},
		{"4 42", "-1"},
		{"100 40021", "5"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve16953(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
