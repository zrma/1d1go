package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayLinkedList(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-linked-list/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	funcPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Println }()

	const (
		want = `2
3
4
1
`
	)

	var arr = []int{2, 3, 4, 1}

	displayLinkedList(arr)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
