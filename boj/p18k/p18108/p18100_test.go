package p18108

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve18108(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18108")

	for i, tt := range []struct {
		given int
		want  int
	}{
		{2541, 1998},
		{543, 0},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve18108(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
