package p4600

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test_data/p4673_want.txt
var p4673want string

func TestSolve4673(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4673")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	Solve4673(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, p4673want, got)
}

func TestDecompose(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want int
	}{
		{1, 2},
		{33, 39},
		{39, 51},
		{51, 57},
		{57, 69},
		{69, 84},
		{84, 96},
		{96, 111},
		{111, 114},
		{114, 120},
		{120, 123},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := decompose(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
