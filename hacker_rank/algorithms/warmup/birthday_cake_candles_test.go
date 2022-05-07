package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthdayCakeCandles(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/birthday-cake-candles/problem")

	arr := []int32{3, 2, 1, 3}
	const want = 2

	got := birthdayCakeCandles(arr)
	assert.EqualValues(t, want, got)
}
