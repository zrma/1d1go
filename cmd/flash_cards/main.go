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

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fprintf(writer, "The term for card #%d:\n", i)
		scanner.Scan()
		label := scanner.Text()

		_, _ = fmt.Fprintf(writer, "The definition for card #%d:\n", i)
		scanner.Scan()
		answer := scanner.Text()

		cards[i-1] = card{label, answer}
	}

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fprintf(writer, "Print the definition of \"%s\":\n", cards[i-1].label)
		scanner.Scan()
		answer := scanner.Text()

		if answer == cards[i-1].answer {
			_, _ = fmt.Fprintln(writer, "Correct!")
		} else {
			_, _ = fmt.Fprintf(writer, "Wrong. The right answer is \"%s\".\n", cards[i-1].answer)
		}
	}
}
