package arrays

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestArrayManipulation(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/crush/problem")

	for i, tt := range []struct {
		n       int
		queries [][]int
		want    int64
	}{
		{
			n:       10,
			queries: [][]int{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}},
			want:    10,
		},
		{
			n:       5,
			queries: [][]int{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}},
			want:    200,
		},
		{
			n:       10,
			queries: [][]int{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}},
			want:    31,
		},
		{
			n:       10,
			queries: [][]int{{3, 5, 7}, {2, 6, 8}, {1, 8, 1}, {5, 9, 15}},
			want:    31,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := arrayManipulation(tt.n, tt.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestArrayManipulationPerformance(t *testing.T) {
	file, err := os.Open("./test_data/array_manipulation.csv")
	assert.NoError(t, err)
	defer func() {
		err := file.Close()
		assert.NoError(t, err)
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	assert.NoError(t, err)

	var arr [][]int
	for _, row := range rows {
		begin, err := strconv.Atoi(strings.TrimSpace(row[0]))
		assert.NoError(t, err)

		end, err := strconv.Atoi(strings.TrimSpace(row[1]))
		assert.NoError(t, err)

		value, err := strconv.Atoi(strings.TrimSpace(row[2]))
		assert.NoError(t, err)

		arr = append(arr, []int{begin, end, value})
	}

	assert.Len(t, arr, 100000)
	const want = 2506721627

	assert.Eventually(t, func() bool {
		got := arrayManipulation(10000000, arr)
		return assert.EqualValues(t, want, got)
	}, time.Second*3, time.Millisecond*100, "시간 초과")
}
