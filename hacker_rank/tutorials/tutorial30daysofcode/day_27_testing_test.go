package tutorial30daysofcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumIndex(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-testing/problem")

	t.Run("emptyArray function returns an empty slice just fine.", func(t *testing.T) {
		const want = invalidIdx
		arr := emptyArray()

		got, err := minimumIndex(arr)
		assert.Error(t, err)
		assert.EqualError(t, err, errMsg)
		assert.Equal(t, want, got)
	})

	t.Run("uniqueValues() function returns a slice containing two or more unique values.", func(t *testing.T) {
		seq := uniqueValues()
		assert.GreaterOrEqual(t, len(seq), 2)

		m := make(map[int]interface{})
		for _, num := range seq {
			_, ok := m[num]
			assert.False(t, ok)
			m[num] = nil
		}

		got, err := minimumIndex(seq)
		assert.NoError(t, err)
		assert.Equal(t, 0, got, "해당 풀이에서는 minimumIndex = 0")
	})

	t.Run("exactlyTwoDifferentMinimums() function returns a slice of the two smallest integers and any other larger integer of zero or more.", func(t *testing.T) {
		seq := exactlyTwoDifferentMinimums()
		assert.GreaterOrEqual(t, len(seq), 2)

		tmp := make([]int, len(seq))
		copy(tmp, seq)
		sort.Slice(tmp, func(i, j int) bool {
			return i > j
		})
		assert.Equal(t, tmp[0], tmp[1])
		assert.True(t, len(tmp) == 2 || tmp[1] < tmp[2])

		got, err := minimumIndex(seq)
		assert.NoError(t, err)
		assert.Equal(t, 1, got, "해당 풀이에서는 minimumIndex = 1")
	})
}
