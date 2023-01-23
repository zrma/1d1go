package p11800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11866(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11866")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"7 3",
			"<3, 6, 2, 7, 5, 1, 4>",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11866(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
