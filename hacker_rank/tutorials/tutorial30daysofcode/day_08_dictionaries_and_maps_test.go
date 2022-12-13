package tutorial30daysofcode

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionariesAndMaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem")

	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	funcPrintf = func(format string, a ...any) (n int, err error) {
		return fmt.Fprintf(writer, format, a...)
	}
	defer func() { funcPrintf = fmt.Printf }()
	funcPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Println }()

	const (
		want = `sam=99912222
Not found
harry=12299933
`
	)

	dictionariesAndMaps(3, []string{
		"sam 99912222",
		"tom 11122222",
		"harry 12299933",
		"sam",
		"edward",
		"harry",
	})

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
