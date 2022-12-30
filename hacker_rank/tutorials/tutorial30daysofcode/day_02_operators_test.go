package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperators(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-operators/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	funcPrint = func(a ...any) (n int, err error) {
		return fmt.Fprint(writer, a...)
	}
	defer func() { funcPrint = fmt.Print }()

	const want = "15"

	operators(12.00, 20, 8)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
