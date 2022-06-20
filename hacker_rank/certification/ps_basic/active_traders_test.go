package ps_basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActiveTraders(t *testing.T) {
	t.Log("https://www.hackerrank.com/skills-verification/problem_solving_basic")

	for i, tt := range []struct {
		give []string
		want []string
	}{
		{
			give: []string{
				"Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha",
				"Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Beta",
			},
			want: []string{"Alpha", "Beta", "Omega"},
		},
		{
			give: []string{
				"Alpha", "Beta", "Zeta", "Beta", "Zeta", "Zeta", "Epsilon", "Beta", "Zeta", "Beta",
				"Zeta", "Beta", "Delta", "Zeta", "Beta", "Zeta", "Beta", "Zeta", "Beta", "Zeta",
				"Beta",
			},
			want: []string{"Beta", "Zeta"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := mostActive(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
