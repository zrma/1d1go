package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetsReview(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-loops/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	fmtPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { fmtPrintln = fmt.Println }()

	const (
		want = `Hce akr
Rn ak
`
	)

	letsReview("Hacker")
	letsReview("Rank")

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
