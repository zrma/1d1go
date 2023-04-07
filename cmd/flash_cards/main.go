package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	FlashCards(scanner, os.Stdout)
}

func FlashCards(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	//label := scanner.Text()

	scanner.Scan()
	question := scanner.Text()

	scanner.Scan()
	answer := scanner.Text()

	if question == answer {
		_, _ = fmt.Fprintln(writer, "Your answer is right!")
	} else {
		_, _ = fmt.Fprintln(writer, "Your answer is wrong...")
	}
}
