package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryNumbers(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-binary-numbers/problem")

	for _, tt := range []struct {
		n    int32
		want string
	}{
		{5, "1"},
		{13, "2"},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			fmtPrint = func(a ...any) (n int, err error) {
				return fmt.Fprint(writer, a...)
			}
			defer func() { fmtPrint = fmt.Print }()

			binaryNumbers(tt.n)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
