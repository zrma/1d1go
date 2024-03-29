package day1

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusMinus(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/plus-minus/problem")
	t.Log("https://www.hackerrank.com/challenges/one-week-preparation-kit-plus-minus/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	fmtPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { fmtPrintf = fmt.Printf }()

	const (
		want = `0.500000
0.333333
0.166667
`
	)

	plusMinus([]int32{-4, 3, -9, 0, 4, 1})

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
