package p1000

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1008(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1008")

	for i, tt := range []struct {
		a, b int
		want float64
	}{
		{1, 3, 1.0 / 3.0},
		{1, 3, 0.33333333333333333333333333333333},
		{4, 5, 4.0 / 5.0},
		{4, 5, 0.8},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			const epsilon = 1e-9

			got := Solve1008(tt.a, tt.b)
			assert.InDelta(t, tt.want, got, epsilon)
			assert.InEpsilon(t, tt.want, got, epsilon)
		})
	}
}
