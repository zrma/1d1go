package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestStringWriter(t *testing.T) {
	writer := utils.NewStringWriter()

	{
		n, err := fmt.Fprintln(writer, "Hello")
		assert.NoError(t, err)
		assert.Equal(t, 5+1, n)
	}
	{
		n, err := fmt.Fprintln(writer, "World!")
		assert.NoError(t, err)
		assert.Equal(t, 6+1, n)
	}
	{
		n, err := fmt.Fprintf(writer, "%d + %d = %d\n", 1, 2, 1+2)
		assert.NoError(t, err)
		assert.Equal(t, 10, n)
	}

	{
		got := writer.String()
		assert.Empty(t, got)
	}
	{
		err := writer.Flush()
		assert.NoError(t, err)

		got := writer.String()
		const want = `Hello
World!
1 + 2 = 3
`
		assert.Equal(t, want, got)
	}
}
