package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleArraySum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/simple-array-sum/problem")

	arr := []int32{1, 2, 3, 4, 10, 11}
	const want = 31

	got := simpleArraySum(arr)
	assert.EqualValues(t, want, got)
}
