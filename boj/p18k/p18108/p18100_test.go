package p18108_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18108"
)

func TestSolve18108(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18108")

	for i, tt := range []struct {
		year int
		want int
	}{
		{2541, 1998},
		{543, 0},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := p18108.Solve18108(tt.year)
			assert.Equal(t, tt.want, got)
		})
	}
}
