package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestDataType(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-data-types/problem")

	writer := utils.NewStringWriter()
	funcPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { funcPrintf = fmt.Printf }()
	funcPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Println }()

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

	got := writer.String()
	assert.Equal(t, want, got)
}
