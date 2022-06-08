package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestMoreException(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-more-exceptions/problem")

	for _, tt := range []struct {
		n, p int64
		want string
	}{
		{3, 5, "243"},
		{2, 4, "16"},
		{-1, -2, "n and p should be non-negative"},
		{-1, 3, "n and p should be non-negative"},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.n, tt.p), func(t *testing.T) {
			writer := utils.NewStringWriter()
			funcPrint = func(a ...any) (n int, err error) {
				return fmt.Fprint(writer, a...)
			}
			defer func() { funcPrint = fmt.Print }()

			moreException(tt.n, tt.p)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
