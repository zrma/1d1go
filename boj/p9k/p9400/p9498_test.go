package p9400

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9498(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9498")

	for _, tt := range []struct {
		score int
		want  string
	}{
		{100, "A"},
		{90, "A"},
		{89, "B"},
		{80, "B"},
		{79, "C"},
		{70, "C"},
		{69, "D"},
		{60, "D"},
		{59, "F"},
		{0, "F"},
	} {
		t.Run(fmt.Sprintf("%d", tt.score), func(t *testing.T) {
			got := Solve9498(tt.score)
			assert.Equal(t, tt.want, got)
		})
	}
}
