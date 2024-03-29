package p1000

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1008(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1008")

	tests := []struct {
		give string
		want string
	}{
		{"1 3", "0.33333333333333333333333333333333"},
		{"4 5", "0.8"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1008(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			want, err := strconv.ParseFloat(tt.want, 64)
			assert.NoError(t, err)

			writen := buf.String()
			got, err := strconv.ParseFloat(writen, 64)
			assert.NoError(t, err)

			const epsilon = 1e-6
			assert.InDelta(t, want, got, epsilon)
		})
	}
}
