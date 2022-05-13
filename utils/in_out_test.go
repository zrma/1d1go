package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIO(t *testing.T) {
	var (
		inCalled  bool
		outCalled bool
		inArgs    = []any{1, 2, 3}
		outArgs   = []any{4, 5, 6}
	)

	const (
		wantN0 = 123
		wantN1 = 456
	)

	fnIn = func(args ...any) (n int, err error) {
		inCalled = true
		assert.Len(t, args, len(inArgs))
		assert.Equal(t, inArgs, args)
		return wantN0, nil
	}
	defer func() { fnIn = fmt.Scan }()

	fnOut = func(args ...any) (n int, err error) {
		outCalled = true
		assert.Len(t, args, len(outArgs))
		assert.Equal(t, outArgs, args)
		return wantN1, nil
	}

	io := NewIO()
	n, err := io.Scan(inArgs...)
	assert.NoError(t, err)
	assert.Equal(t, wantN0, n)
	assert.True(t, inCalled)

	n, err = io.Println(outArgs...)
	assert.NoError(t, err)
	assert.Equal(t, wantN1, n)
	assert.True(t, outCalled)
}
