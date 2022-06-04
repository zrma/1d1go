package utils

import (
	"bufio"
	"os"
	"strings"

	"go.uber.org/multierr"
)

func ReadStringFromFile(fileName string) (s string, e error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func() {
		if e != nil {
			e = multierr.Append(e, file.Close())
		}
	}()

	r := bufio.NewScanner(bufio.NewReader(file))
	r.Split(bufio.ScanLines)

	var lines []string
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	s = strings.Join(lines, "\n") + "\n"
	e = r.Err()
	return s, e
}
