package utils_test

import (
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestNewStringScanner(t *testing.T) {
	const s = `3
3 123 456 789
4 12 345 67 890
5 12 23 56 78
`
	scanner := utils.NewStringScanner(s)
	assert.NoError(t, scanner.Err())

	assert.True(t, scanner.Scan())
	assert.Equal(t, "3", scanner.Text())

	for _, tt := range []struct {
		line int
		nums []int
	}{
		{3, []int{123, 456, 789}},
		{4, []int{12, 345, 67, 890}},
		{5, []int{12, 23, 56, 78}},
	} {
		assert.True(t, scanner.Scan())
		line, err := strconv.Atoi(scanner.Text())
		assert.NoError(t, err)
		assert.Equal(t, tt.line, line)
		for _, num := range tt.nums {
			assert.True(t, scanner.Scan())
			got, err := strconv.Atoi(scanner.Text())
			assert.NoError(t, err)
			assert.Equal(t, num, got)
		}
	}

	assert.False(t, scanner.Scan())
	err := scanner.Err()
	assert.Error(t, err)
	assert.Equal(t, io.EOF, err)
}
