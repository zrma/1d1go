package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-regex-patterns/problem")

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	funcPrintln = func(a ...any) (n int, err error) {
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

	got := buf.String()
	assert.Equal(t, want, got)
}
