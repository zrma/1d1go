package p11700_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11700"
)

func TestSolve11726(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11726")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"2", "2"},
		{"9", "55"},
		{"12", "233"},
		{"1000", "1115"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p11700.Solve11726(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
