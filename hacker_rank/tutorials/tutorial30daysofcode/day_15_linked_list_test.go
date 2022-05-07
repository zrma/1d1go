package tutorial30daysofcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestDisplayLinkedList(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-linked-list/problem")

	var arr = []int{2, 3, 4, 1}
	want := make([]string, len(arr))
	for i, n := range arr {
		want[i] = strconv.Itoa(n)
	}
	got, err := utils.GetPrinted(func() {
		displayLinkedList(arr)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
