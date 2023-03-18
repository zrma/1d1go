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

		switch s {
		case "/help":
			_, _ = fmt.Fprintln(writer, "The program calculates the sum of numbers")
		case "/exit":
			_, _ = fmt.Fprintln(writer, "Bye!")
			return
		default:
			process(s, writer)
		}
	}
}

func process(input string, writer io.Writer) {
	ss := strings.Split(input, " ")
	tokens := trimTokens(ss)

	if len(tokens) == 0 {
		return
	}

	res := calc(tokens)
	_, _ = fmt.Fprintln(writer, res)
}

func trimTokens(token []string) []string {
	var res []string
	for _, s := range token {
		s = strings.Trim(s, " ")
		if s == "" {
			continue
		}

		if isOperator(s) {
			s = pruneOperator(s)
		}
		res = append(res, s)
	}
	return res
}

func pruneOperator(s string) string {
	res := s
	for len(res) > 1 {
		if res[0] == res[1] {
			res = "+" + res[2:]
		} else {
			res = "-" + res[2:]
		}
	}
	return res
}

func calc(tokens []string) int {
	var res int
	var op string

	for _, token := range tokens {
		if isOperator(token) {
			op = token
			continue
		}

		n, err := strconv.Atoi(token)
		if err != nil {
			continue
		}

		switch op {
		case "+":
			res += n
		case "-":
			res -= n
		default:
			res = n
		}
	}

	return res
}

func isOperator(token string) bool {
	return !strings.ContainsAny(token, "0123456789")
}
