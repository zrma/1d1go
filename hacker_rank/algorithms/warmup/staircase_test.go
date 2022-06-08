package warmup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestStaircase(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/staircase/problem")

	writer := utils.NewStringWriter()
	funcPrintln = func(a ...any) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Println }()

	const (
		want = `     #
    ##
   ###
  ####
 #####
######
`
	)

	staircase(6)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
