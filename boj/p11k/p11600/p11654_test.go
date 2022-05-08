package p11600

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			got := Solve11654(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
