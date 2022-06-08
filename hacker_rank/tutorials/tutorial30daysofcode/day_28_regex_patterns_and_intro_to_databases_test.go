package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestFilter(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-regex-patterns/problem")

	writer := utils.NewStringWriter()
	funcPrintln = func(a ...interface{}) (n int, err error) {
		return fmt.Fprintln(writer, a...)
	}
	defer func() { funcPrintln = fmt.Println }()

	//noinspection SpellCheckingInspection
	const (
		want = `julia
julia
riya
samantha
tanya
`
	)
	//noinspection SpellCheckingInspection
	filter([][]string{
		{"riya", "riya@gmail.com"},
		{"julia", "julia@julia.me"},
		{"julia", "sjulia@gmail.com"},
		{"julia", "julia@gmail.com"},
		{"samantha", "samantha@gmail.com"},
		{"tanya", "tanya@gmail.com"},
	})

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
