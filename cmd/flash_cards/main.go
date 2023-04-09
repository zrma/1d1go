package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	FlashCards(scanner, os.Stdout)
}

func FlashCards(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	_, _ = fmt.Fprintln(writer, "Input the number of cards:")

	scanner.Scan()
	v := scanner.Text()
	n, _ := strconv.Atoi(v)

	type card struct {
		label  string
		answer string
	}

	cards := make([]card, n)
	labels := make(map[string]int)
	answers := make(map[string]int)

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "The term for card #%d:\n", i+1)

		var label string
		for scanner.Scan() {
			label = scanner.Text()

			if _, ok := labels[label]; !ok {
				break
			}

			_, _ = fmt.Fprintf(writer, "The term \"%s\" already exists. Try again:\n", label)
		}

		_, _ = fmt.Fprintf(writer, "The definition for card #%d:\n", i+1)

		var answer string

		for scanner.Scan() {
			answer = scanner.Text()

			if _, ok := answers[answer]; !ok {
				break
			}

			_, _ = fmt.Fprintf(writer, "The definition \"%s\" already exists. Try again:\n", answer)
		}

		cards[i] = card{label, answer}
		labels[label] = i
		answers[answer] = i
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		_, _ = fmt.Fprintf(writer, "Print the definition of \"%s\":\n", cards[i].label)

		answer := scanner.Text()

		if answer == cards[i].answer {
			_, _ = fmt.Fprintln(writer, "Correct!")
			continue
		}

		_, ok := answers[answer]
		if !ok {
			_, _ = fmt.Fprintf(writer, "Wrong. The right answer is \"%s\".\n", cards[i].answer)
			continue
		}

		_, _ = fmt.Fprintf(writer, "Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", cards[i].answer, cards[answers[answer]].label)
	}
}
