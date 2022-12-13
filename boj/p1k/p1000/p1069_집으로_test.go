package p1000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
)

func TestSolve1069(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1069")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"6 8 5 3", "6.0"},
		{"3 4 6 3", "4.0"},
		{"318 445 1200 800", "546.9451526432975"},
		{"400 300 150 10", "40.0"},
		{"6 8 3 2", "7.0"},
		{"10 10 1000 5", "10.0"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1000.Solve1069(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			want, err := strconv.ParseFloat(tt.want, 64)
			assert.NoError(t, err)

			writen := buf.String()
			got, err := strconv.ParseFloat(writen, 64)
			assert.NoError(t, err)

			const epsilon = 1e-9
			assert.InDelta(t, want, got, epsilon)
		})
	}
}
