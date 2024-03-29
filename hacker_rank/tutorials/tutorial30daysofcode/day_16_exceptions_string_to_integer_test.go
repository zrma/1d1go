package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInteger(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	fmtPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { fmtPrintln = fmt.Print }()

	const (
		want = `3
Bad String
`
	)

	stringToInteger("3")
	stringToInteger("za")

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
