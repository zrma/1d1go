package p2000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolve2004(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2004")

	tests := []struct {
		give string
		want string
	}{
		{"25 12", "2"},
		{"5 1", "0"},
		{"5 4", "0"},
		{"1 1", "0"},
		{"10 2", "0"},
		{"2000000000 1000000000", "1"},
		{"2000000000 1", "9"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2004(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSolve2004_Performance(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"2000000000 1", "9"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			assert.Eventually(t, func() bool {
				Solve2004(reader, writer)

				err := writer.Flush()
				assert.NoError(t, err)

				got := buf.String()
				return assert.Equal(t, tt.want, got)
			}, time.Second, time.Millisecond*100, "시간 초과")
		})
	}
}
