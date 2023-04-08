package p1200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1251(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1251")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{"mobitel", "bometil"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1251(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
