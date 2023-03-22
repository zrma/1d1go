package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	SmartCalc(scanner, os.Stdout)
}

func SmartCalc(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()

		if strings.HasPrefix(s, "/") {
			switch s {
			case "/help":
				_, _ = fmt.Fprintln(writer, "The program calculates the sum of numbers")
			case "/exit":
				_, _ = fmt.Fprintln(writer, "Bye!")
				return
			default:
				_, _ = fmt.Fprintln(writer, "Unknown command")
			}
		} else {
			process(s, writer)
		}
	}
}

// process implements calculator logic, if invalid input, print "Invalid expression"
func process(input string, writer io.Writer) {
	tokens := strings.Split(input, " ")

	// remove empty tokens
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "" {
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}

	if len(tokens) == 0 {
		return
	}

	var res int
	var err error
	var minus bool
	for i, token := range tokens {
		if i%2 == 0 {
			// number
			var n int
			n, err = strconv.Atoi(token)
			if err != nil {
				break
			}
			if minus {
				res -= n
			} else {
				res += n
			}
		} else {
			// operator
			op, err0 := pruneOperator(token)
			if err0 != nil {
				err = err0
				break
			}
			minus = op == "-"
		}
	}

	if err != nil {
		_, _ = fmt.Fprintln(writer, "Invalid expression")
	} else {
		_, _ = fmt.Fprintln(writer, res)
	}
}

func pruneOperator(s string) (string, error) {
	res := "+"
	for _, c := range s {
		switch c {
		case '+':
		case '-':
			if res == "+" {
				res = "-"
			} else {
				res = "+"
			}
		default:
			return "", fmt.Errorf("invalid operator: %c", c)
		}
	}
	return res, nil
}
