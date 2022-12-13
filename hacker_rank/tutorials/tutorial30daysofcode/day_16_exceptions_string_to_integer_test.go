package tutorial30daysofcode

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInteger(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	funcPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Print }()

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
