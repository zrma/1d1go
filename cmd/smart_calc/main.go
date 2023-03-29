package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	SmartCalc(scanner, os.Stdout)
}

func SmartCalc(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	variables := make(map[string]int)

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
		} else if strings.Contains(s, "=") {
			assignment(s, variables, writer)
		} else {
			process(s, variables, writer)
		}
	}
}

func assignment(input string, variables map[string]int, writer io.Writer) {
	input = strings.TrimSpace(input)

	tokens := strings.Split(input, "=")
	if len(tokens) != 2 {
		_, _ = fmt.Fprintln(writer, "Invalid assignment")
		return
	}

	variable := strings.TrimSpace(tokens[0])
	value := strings.TrimSpace(tokens[1])

	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(variable) {
		_, _ = fmt.Fprintln(writer, "Invalid identifier")
		return
	}

	if regex.MatchString(value) {
		// variable
		if v, ok := variables[value]; ok {
			variables[variable] = v
		} else {
			_, _ = fmt.Fprintln(writer, "Unknown variable")
		}
	} else {
		// number
		n, err := strconv.Atoi(value)
		if err != nil {
			_, _ = fmt.Fprintln(writer, "Invalid assignment")
			return
		}
		variables[variable] = n
	}
}

// process implements calculator logic, if invalid input, print "Invalid expression"
func process(input string, variables map[string]int, writer io.Writer) {
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
				// variable
				if v, ok := variables[token]; ok {
					n = v
					err = nil
				} else {
					regex := regexp.MustCompile(`^[a-zA-Z]+$`)
					if regex.MatchString(token) {
						_, _ = fmt.Fprintln(writer, "Unknown variable")
						return
					}

					//goland:noinspection SpellCheckingInspection
					const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

					if strings.ContainsAny(token, alphabet) {
						_, _ = fmt.Fprintln(writer, "Invalid identifier")
						return
					}
					break
				}
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
