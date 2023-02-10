package p2700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2754(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2754")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"A+", "4.3"},
		{"A0", "4.0"},
		{"A-", "3.7"},
		{"B+", "3.3"},
		{"B0", "3.0"},
		{"B-", "2.7"},
		{"C+", "2.3"},
		{"C0", "2.0"},
		{"C-", "1.7"},
		{"D+", "1.3"},
		{"D0", "1.0"},
		{"D-", "0.7"},
		{"F", "0.0"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2754(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
