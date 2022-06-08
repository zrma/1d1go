package codesprint5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestExceedingTheSpeedLimit(t *testing.T) {
	for _, tt := range []struct {
		n    int32
		want string
	}{
		{
			100, "3000 Warning"},
		{
			140, "25000 License removed"},
		{
			85, "0 No punishment"},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			writer := utils.NewStringWriter()
			funcPrintf = func(format string, a ...any) (n int, err error) {
				return fmt.Fprintf(writer, format, a...)
			}
			defer func() { funcPrintf = fmt.Printf }()

			exceedingTheSpeedLimit(tt.n)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
