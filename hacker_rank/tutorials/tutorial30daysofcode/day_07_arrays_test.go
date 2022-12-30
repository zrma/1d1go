package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintReverse(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-arrays/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	funcPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { funcPrintf = fmt.Printf }()

	const want = "2 3 4 1"

	printReverse([]int32{1, 4, 3, 2})

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
