package p1400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1427(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1427")

	tests := []struct {
		give string
		want string
	}{
		{"2143", "4321"},
		{"999998999", "999999998"},
		{"61423", "64321"},
		{"500613009", "965310000"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1427(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
