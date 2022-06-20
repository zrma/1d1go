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

func TestActivityNotifications(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem")

	for i, tt := range []struct {
		expenditure []int32
		d           int32
		want        int32
	}{
		{
			expenditure: []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
			d:           5,
			want:        2,
		},
		{
			expenditure: []int32{1, 2, 3, 4, 4},
			d:           4,
			want:        0,
		},
		{
			expenditure: []int32{10, 20, 30, 40, 50},
			d:           3,
			want:        1,
		},
		{
			expenditure: []int32{10, 20, 30, 40, 50, 80},
			d:           4,
			want:        2,
		},
	} {
		t.Run(fmt.Sprintf("%d/counting sort", i), func(t *testing.T) {
			got := activityNotifications(tt.expenditure, tt.d)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/sort.Slice", i), func(t *testing.T) {
			got := activityNotificationsWithSort(tt.expenditure, tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestActivityNotificationsPerformance(t *testing.T) {
	file, err := os.Open("./test_data/activity_notifications.csv")
	assert.NoError(t, err)
	defer func() {
		err := file.Close()
		assert.NoError(t, err)
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	assert.NoError(t, err)

	var expenditure []int32
	for _, row := range rows {
		num, err := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 32)
		assert.NoError(t, err)

		expenditure = append(expenditure, int32(num))
	}

	assert.Len(t, expenditure, 200000)

	const (
		d          = 40001
		want int32 = 926
	)

	assert.Eventually(t, func() bool {
		got := activityNotifications(expenditure, d)
		return assert.Equal(t, want, got)
	}, time.Second*3, time.Millisecond*100, "시간 초과")
}

func TestMedianOdd(t *testing.T) {
	for i, tt := range []struct {
		counts []int32
		d      int32
		want   float64
	}{
		{
			[]int32{5, 5, 100, 3},
			7,
			0,
		},
		{
			[]int32{0, 1, 3, 5, 2},
			9,
			3,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := medianOdd(tt.counts, tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMedianEven(t *testing.T) {
	for i, tt := range []struct {
		counts []int32
		d      int32
		want   float64
	}{
		{
			[]int32{5, 5, 100, 3},
			8,
			0,
		},
		{
			[]int32{0, 1, 3, 5, 2},
			8,
			2.5,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := medianEven(tt.counts, tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMedianOddWithSort(t *testing.T) {
	for i, tt := range []struct {
		give []int32
		want float64
	}{
		{
			[]int32{1, 3, 2, 5, 4},
			3,
		},
		{
			[]int32{4, 2, 2, 3, 1},
			2,
		},
		{
			[]int32{1},
			1,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := medianOddWithSort(tt.give, len(tt.give)/2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMedianEvenWithSort(t *testing.T) {
	for i, tt := range []struct {
		give []int32
		want float64
	}{
		{
			[]int32{1, 3, 2, 5, 4, 3},
			3,
		},
		{
			[]int32{4, 2, 2, 3, 1, 4},
			2.5,
		},
		{
			[]int32{1, 2},
			1.5,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := medianEvenWithSort(tt.give, len(tt.give)/2-1, len(tt.give)/2)
			assert.Equal(t, tt.want, got)
		})
	}
}
