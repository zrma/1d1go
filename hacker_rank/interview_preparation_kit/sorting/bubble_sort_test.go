package sorting

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

func TestCountSwaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/ctci-bubble-sort/problem")

	for i, tt := range []struct {
		give []int32
		want string
	}{
		{
			give: []int32{6, 4, 1},
			want: `Array is sorted in 3 swaps.
First Element: 1
Last Element: 6
`,
		},
		{
			give: []int32{1, 2, 3},
			want: `Array is sorted in 0 swaps.
First Element: 1
Last Element: 3
`,
		},
		{
			give: []int32{3, 2, 1},
			want: `Array is sorted in 3 swaps.
First Element: 1
Last Element: 3
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			funcPrintf = func(format string, a ...any) (n int, err error) {
				return fmt.Fprintf(writer, format, a...)
			}
			defer func() { funcPrintf = fmt.Printf }()

			countSwaps(tt.give)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountSwapsPerformance(t *testing.T) {
	file, err := os.Open("./test_data/bubble_sort.csv")
	assert.NoError(t, err)
	defer func() {
		err := file.Close()
		assert.NoError(t, err)
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	assert.NoError(t, err)

	var arr []int32
	for _, row := range rows {
		num, err := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 32)
		assert.NoError(t, err)

		arr = append(arr, int32(num))
	}

	assert.Len(t, arr, 528)

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	funcPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { funcPrintf = fmt.Printf }()

	const (
		want = `Array is sorted in 68472 swaps.
First Element: 4842
Last Element: 1994569
`
	)
	assert.Eventually(t, func() bool {
		countSwaps(arr)

		err := writer.Flush()
		assert.NoError(t, err)

		got := buf.String()
		return assert.Equal(t, want, got)
	}, time.Second, time.Millisecond*100, "시간 초과")
}
