package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataType(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-data-types/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	fmtPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { fmtPrintf = fmt.Printf }()
	fmtPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { fmtPrintln = fmt.Println }()

	const (
		want = `16
8.0
HackerRank is the best place to learn and practice coding!
`
	)

	dataType(
		4, 12,
		4.0, 4.0,
		"HackerRank ", "is the best place to learn and practice coding!",
	)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
