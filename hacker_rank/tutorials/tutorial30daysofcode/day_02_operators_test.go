package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestOperators(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-operators/problem")

	writer := utils.NewStringWriter()
	funcPrint = func(a ...any) (n int, err error) {
		return fmt.Fprint(writer, a...)
	}
	defer func() { funcPrint = fmt.Print }()

	const want = "15"

	operators(12.00, 20, 8)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
