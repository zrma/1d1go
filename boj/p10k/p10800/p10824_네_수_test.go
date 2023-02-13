package p10800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10824(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10824")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"10 20 30 40", "4060"},
		{"1 2 3 4", "46"},
		{"100 200 300 400", "400600"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10824(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
