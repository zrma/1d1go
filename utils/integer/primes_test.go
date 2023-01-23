package integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSieveOfEratosthenes(t *testing.T) {
	t.Log("https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes")

	for i, tt := range []struct {
		give int
		want []bool
	}{
		{
			10,
			[]bool{true, true, false, false, true, false, true, false, true, true},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := BuildSieveOfEratosthenes(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
