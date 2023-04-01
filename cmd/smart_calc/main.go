package main

import (
	"bufio"
	"errors"
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
		} else {
			n, ok, err := process(s, variables, writer)
			if err != nil {
				_, _ = fmt.Fprintln(writer, err.Error())
				continue
			}
			if ok {
				_, _ = fmt.Fprintln(writer, n)
			}
		}
	}
}

func process(input string, variables map[string]int, writer io.Writer) (int, bool, error) {
	if input == "" {
		return 0, false, nil
	}

	if strings.Contains(input, "=") {
		if err := processAssignment(input, variables, writer); err != nil {
			return 0, false, err
		}
		return 0, false, nil
	}

	v, ok, err := processUnary(input, variables)
	if err != nil {
		return 0, false, err
	}
	if ok {
		return v, true, nil
	}

	res, err := processExpression(input, variables)
	if err != nil {
		return 0, false, err
	}
	return res, true, nil
}

//goland:noinspection GoErrorStringFormat
var (
	errInvalidAssignment = errors.New("Invalid assignment")
	errInvalidIdentifier = errors.New("Invalid identifier")
	errInvalidExpression = errors.New("Invalid expression")
	errUnknownVariable   = errors.New("Unknown variable")
	errUnknownOperator   = errors.New("Unknown operator")
)

func processAssignment(input string, variables map[string]int, writer io.Writer) error {
	tokens := strings.Split(input, "=")
	if len(tokens) != 2 {
		return errInvalidAssignment
	}

	lhs := strings.TrimSpace(tokens[0])
	rhs := strings.TrimSpace(tokens[1])

	if !isValidVariable(lhs) {
		return errInvalidIdentifier
	}

	n, err := strconv.Atoi(rhs)
	if err == nil {
		variables[lhs] = n
		return nil
	}

	n, ok, err := process(rhs, variables, writer)
	if err != nil {
		if errors.Is(err, errInvalidIdentifier) {
			return errInvalidAssignment
		}
		return err
	}
	if ok {
		variables[lhs] = n
	}
	return nil
}

func processUnary(input string, variables map[string]int) (int, bool, error) {
	input = strings.TrimSpace(input)

	if v, ok := variables[input]; ok {
		return v, true, nil
	}

	if isValidVariable(input) {
		return 0, false, errUnknownVariable
	}

	n, err := strconv.Atoi(input)
	if err == nil {
		return n, true, nil
	}

	for _, op := range operators {
		if strings.Contains(input, op) {
			return 0, false, nil
		}
	}

	if strings.Contains(input, " ") {
		return 0, false, nil
	}

	return 0, false, errInvalidIdentifier
}

func processExpression(input string, variables map[string]int) (int, error) {
	expr, err := infixToPostfix(input, variables)
	if err != nil {
		return 0, err
	}
	return calcPostfix(expr)
}

var operatorsPriority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

func calcPostfix(queue []string) (int, error) {
	stack := make([]int, 0)

	for len(queue) > 0 {
		token := queue[0]
		queue = queue[1:]

		if _, ok := operatorsPriority[token]; ok {
			if len(stack) < 2 {
				return 0, errInvalidExpression
			}

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			case "^":
				res := 1
				for i := 0; i < b; i++ {
					res *= a
				}
				stack = append(stack, res)
			default:
				return 0, errUnknownOperator
			}
		} else {
			n, err := strconv.Atoi(token)
			if err != nil {
				return 0, errUnknownOperator
			}
			stack = append(stack, n)
		}
	}

	if len(stack) != 1 {
		return 0, errInvalidExpression
	}

	return stack[0], nil
}

func infixToPostfix(input string, variables map[string]int) ([]string, error) {
	stack := make([]string, 0)
	output := make([]string, 0)

	tokens, err := parseExpr(input)
	if err != nil {
		return nil, err
	}

	for len(tokens) > 0 {
		token := tokens[0]
		tokens = tokens[1:]

		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		if token == "(" {
			stack = append(stack, token)
			continue
		}

		if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return nil, errInvalidExpression
			}
			continue
		}

		if _, ok := operatorsPriority[token]; ok {
			if len(stack) == 0 {
				stack = append(stack, token)
				continue
			}

			for {
				if len(stack) == 0 {
					break
				}
				if _, ok := operatorsPriority[stack[len(stack)-1]]; !ok {
					break
				}
				if operatorsPriority[token] > operatorsPriority[stack[len(stack)-1]] {
					break
				}

				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, token)
			continue
		}

		n, err := strconv.Atoi(token)
		if err != nil {
			if !isValidVariable(token) {
				return nil, errInvalidAssignment
			}

			v, ok := variables[token]
			if !ok {
				return nil, errUnknownVariable
			}
			n = v
		}

		output = append(output, strconv.Itoa(n))
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errInvalidExpression
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

var operators = []string{
	"^", "*", "/", "(", ")", "+", "-",
}

func parseExpr(input string) ([]string, error) {
	for _, op := range operators {
		for strings.Contains(input, " "+op) {
			input = strings.ReplaceAll(input, " "+op, op)
		}
		for strings.Contains(input, op+" ") {
			input = strings.ReplaceAll(input, op+" ", op)
		}
	}

	if strings.Contains(input, "**") || strings.Contains(input, "//") {
		return nil, errInvalidExpression
	}

	input = removeDuplicatedOps(input)

	for _, op := range operators {
		input = strings.ReplaceAll(input, op, " "+op+" ")
	}
	return strings.Split(input, " "), nil
}

func removeDuplicatedOps(input string) string {
	s := input
	for strings.Contains(s, "++") || strings.Contains(s, "--") || strings.Contains(s, "+-") || strings.Contains(s, "-+") {
		s = strings.ReplaceAll(s, "++", "+")
		s = strings.ReplaceAll(s, "--", "+")
		s = strings.ReplaceAll(s, "+-", "-")
		s = strings.ReplaceAll(s, "-+", "-")
	}
	return s
}

func isValidVariable(s string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	return regex.MatchString(s)
}
