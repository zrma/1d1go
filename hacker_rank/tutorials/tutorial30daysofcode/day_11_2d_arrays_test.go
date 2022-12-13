package tutorial30daysofcode

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourGlassSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-2d-arrays/problem")

	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	funcPrint = func(a ...any) (n int, err error) {
		return fmt.Fprint(writer, a...)
	}
	defer func() { funcPrint = fmt.Print }()

	const want = "19"

	arr2D := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	hourGlassSum(arr2D)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
