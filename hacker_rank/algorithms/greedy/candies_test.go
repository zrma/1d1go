package greedy

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCanDies(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/candies/problem")

	for _, tt := range []struct {
		description string
		arr         []int32
		want        int64
	}{
		{
			description: "test_0",
			arr:         []int32{1, 2, 2},
			want:        4,
		},
		{
			description: "test_1",
			arr:         []int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1},
			want:        19,
		},
		{
			description: "test_2",
			arr:         []int32{2, 4, 3, 5, 2, 6, 4, 5},
			want:        12,
		},
		{
			description: "test_3",
			arr:         []int32{1, 3, 3, 3, 2, 1},
			want:        10,
		},
		{
			description: "test_4",
			arr:         []int32{1, 3, 3, 3, 2, 2, 2, 3, 3, 3, 1},
			want:        15,
		},
	} {
		got := candies(int32(len(tt.arr)), tt.arr)
		assert.Equal(t, tt.want, got)
	}
}

func TestCandiesPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		file, err := os.Open("./test_data/candies.csv")
		assert.NoError(t, err)
		defer func() {
			err := file.Close()
			assert.NoError(t, err)
		}()

		r := csv.NewReader(bufio.NewReader(file))
		rows, err := r.ReadAll()
		assert.NoError(t, err)

		var arr []int32
		for _, col := range rows {
			num, err := strconv.ParseInt(strings.TrimSpace(col[0]), 10, 32)
			assert.NoError(t, err)

			arr = append(arr, int32(num))
		}

		assert.Len(t, arr, 100000)
		const want = 160929

		got := candies(int32(len(arr)), arr)
		assert.EqualValues(t, want, got)

		return true
	}, time.Second, time.Millisecond*100, "시간 초과")
}
