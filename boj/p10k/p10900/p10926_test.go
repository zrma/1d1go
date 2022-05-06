package p10900

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10926(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10926")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		given string
		want  string
	}{
		{"joonas", "joonas??!"},
		{"baekjoon", "baekjoon??!"},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := Solve10926(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
