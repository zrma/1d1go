package p4600_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4600"
	"1d1go/utils"
)

func TestSolve4673(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4673")

	want, err := readFile(t, "./test_data/p4673_want.txt")
	assert.NoError(t, err)

	writer := utils.NewStringWriter()
	p4600.Solve4673(writer)

	err = writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}

func readFile(tb testing.TB, fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func() {
		err := file.Close()
		assert.NoError(tb, err)
	}()

	r := bufio.NewScanner(bufio.NewReader(file))
	r.Split(bufio.ScanLines)

	var lines []string
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	return strings.Join(lines, "\n") + "\n", r.Err()
}
