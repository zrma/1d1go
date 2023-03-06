package p1200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1212(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1212")

	tests := []struct {
		give string
		want string
	}{
		{"314", "11001100"},
		{"7", "111"},
		{"66", "110110"},
		{"234", "10011100"},
		{"5", "101"},
		{"4", "100"},
		{"1", "1"},
		{"0", "0"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1212(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToBinary(t *testing.T) {
	tests := []struct {
		give byte
		want string
	}{
		{'0', "000"},
		{'1', "001"},
		{'2', "010"},
		{'3', "011"},
		{'4', "100"},
		{'5', "101"},
		{'6', "110"},
		{'7', "111"},
		{'8', ""},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := toBinary(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
