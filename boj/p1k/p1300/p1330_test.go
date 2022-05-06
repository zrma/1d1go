package p1300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1330(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1330")

	for i, tt := range []struct {
		a, b int
		want string
	}{
		{1, 2, "<"},
		{10, 2, ">"},
		{5, 5, "=="},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve1330(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
