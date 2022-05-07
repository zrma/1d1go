package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseArray(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/arrays-ds/problem")

	var (
		arr  = []int32{1, 4, 3, 2}
		want = []int32{2, 3, 4, 1}
	)

	got := reverseArray(arr)
	assert.Equal(t, want, got)
}
