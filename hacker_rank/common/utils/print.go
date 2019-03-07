package utils

import (
	"errors"
	"github.com/go-test/deep"
	"io/ioutil"
	"os"
	"strings"
)

func PrintTest(f func() error, expected []string) error {
	r, w, _ := os.Pipe()
	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
	}()
	os.Stdout = w
	ch := make(chan error)
	go func() {
		ch <- f()
		w.Close()
	}()
	e := <-ch
	stdout, err := ioutil.ReadAll(r)
	if err != nil || e != nil {
		return err
	}
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
