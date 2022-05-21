package p11600_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11600"
)

func TestSolve11654(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11654")

	for _, tt := range []struct {
		s    string
		want int
	}{
		{"A", 65},
		{"B", 66},
		{"C", 67},
		{"0", 48},
		{"1", 49},
		{"9", 57},
		{"a", 97},
		{"z", 122},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := p11600.Solve11654(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
