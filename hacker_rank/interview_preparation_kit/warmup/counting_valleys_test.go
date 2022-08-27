package warmup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingValleys(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/counting-valleys/problem")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want int32
	}{
		{"UDDDUDUU", 1},
		{"", 0},
		{"ABC", 0},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := countingValleys(int32(len(tt.give)), tt.give)
			assert.Equal(t, tt.want, got)
		})
	}

	assert.NotPanics(t, func() {
		const (
			give = "ABC"
			want = 0
		)
		got := countingValleys(int32(len(give))+123, give)
		assert.EqualValues(t, want, got)
	})
}
