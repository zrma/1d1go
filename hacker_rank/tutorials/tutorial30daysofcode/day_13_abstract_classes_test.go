package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractClasses(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-abstract-classes/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	fmtPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { fmtPrintln = fmt.Println }()

	//noinspection SpellCheckingInspection
	const (
		title  = "The Alchemist"
		author = "Paulo Coelho"
		price  = 248
		want   = `Title: The Alchemist
Author: Paulo Coelho
Price: 248
`
	)

	abstractClasses(title, author, price)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
