package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestPrintReverse(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-arrays/problem")

	writer := utils.NewStringWriter()
	funcPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { funcPrintf = fmt.Printf }()

	const want = "2 3 4 1"

	printReverse([]int32{1, 4, 3, 2})

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
