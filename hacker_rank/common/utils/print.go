package utils

import (
	"errors"
	"github.com/go-test/deep"
	"io/ioutil"
	"os"
	"strings"
)

func PrintTest(f func(), expected []string) error {
	r, w, _ := os.Pipe()
	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
		r.Close()
	}()
	os.Stdout = w
	go func() {
		defer w.Close()
		f()
	}()
	stdout, _ := ioutil.ReadAll(r)

	actual := strings.Split(string(stdout), "\n")
	if actual[len(actual)-1] == "" {
		actual = actual[:len(actual)-1]
	}
	diff := deep.Equal(actual, expected)
	if diff != nil {
		return errors.New(strings.Join(diff, "\n"))
	}
	return nil
}
