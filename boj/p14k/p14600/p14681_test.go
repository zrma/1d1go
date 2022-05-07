package p14600

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve14681(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14681")

	for i, tt := range []struct {
		x, y int
		want int
	}{
		{12, 5, 1},
		{-12, 5, 2},
		{-9, -13, 3},
		{9, -13, 4},
		{0, 0, -1},
		{0, -1, -1},
		{369, 0, -1},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve14681(tt.x, tt.y)
			assert.Equal(t, tt.want, got)
		})
	}
}
