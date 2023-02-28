package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2864(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2864")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"11 25", "36 37"},
		{"1430 4862", "6282 6292"},
		{"16796 58786", "74580 85582"},
		{"5 5", "10 12"},
		{"6 6", "10 12"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2864(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
