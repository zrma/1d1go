package p2500

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2588(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2588")

	for i, tt := range []struct {
		a, b int
		want [4]int
	}{
		{
			472, 385,
			[4]int{2360, 3776, 1416, 181720},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve2588(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
