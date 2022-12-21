package p1900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
)

func TestSolve1918(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1918")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"A*(B+C)", "ABC+*"},
		{"A+B", "AB+"},
		{"A+B*C", "ABC*+"},
		{"A+B*C-D/E", "ABC*+DE/-"},
		{"A+B*C-D/E*F", "ABC*+DE/F*-"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1900.Solve1918(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
