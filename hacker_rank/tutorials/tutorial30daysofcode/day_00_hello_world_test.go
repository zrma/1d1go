package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-hello-world/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	fmtPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { fmtPrintln = fmt.Println }()

	const (
		want = `Hello, World.
Welcome to 30 Days of Code!
`
	)

	helloWorld("Welcome to 30 Days of Code!")

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
