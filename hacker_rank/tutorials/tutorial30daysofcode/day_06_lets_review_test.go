package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestLetsReview(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-loops/problem")

	writer := utils.NewStringWriter()
	funcPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Println }()

	const (
		want = `Hce akr
Rn ak
`
	)

	letsReview("Hacker")
	letsReview("Rank")

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
