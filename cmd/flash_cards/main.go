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

	_, _ = fmt.Fprintln(writer, "Card:")
	_, _ = fmt.Fprintln(writer, "purchase")
	_, _ = fmt.Fprintln(writer, "Definition:")
	_, _ = fmt.Fprintln(writer, "buy")
}
