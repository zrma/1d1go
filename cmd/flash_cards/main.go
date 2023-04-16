package main

import (
	"bufio"
	"encoding/json"
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
	cards := newCards()

	for {
		_, _ = fmt.Fprintln(writer, "Input the action (add, remove, import, export, ask, exit):")

		if !scanner.Scan() {
			return
		}
		action := scanner.Text()

		switch action {
		case "add":
			addCard(scanner, writer, cards)
		case "remove":
			removeCard(scanner, writer, cards)
		case "import":
			importCards(scanner, writer, cards)
		case "export":
			exportCards(scanner, writer, cards)
		case "ask":
			askCards(scanner, writer, cards)
		case "exit":
			_, _ = fmt.Fprintln(writer, "Bye bye!")
			return
		}

		_, _ = fmt.Fprintln(writer)
	}
}

func importCards(scanner *bufio.Scanner, writer io.Writer, cards *cardEntries) {
	_, _ = fmt.Fprintln(writer, "File name:")
	scanner.Scan()
	fileName := scanner.Text()

	file, err := os.Open(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(writer, "File not found.\n")
		return
	}
	defer func() { _ = file.Close() }()

	var newCards []CardEntry
	if err := json.NewDecoder(file).Decode(&newCards); err != nil {
		_, _ = fmt.Fprintf(writer, "%+v\n", err)
		return
	}

	for _, card := range newCards {
		cards.upsertCard(card)
	}
	_, _ = fmt.Fprintf(writer, "%d cards have been loaded.\n", len(newCards))
}

func exportCards(scanner *bufio.Scanner, writer io.Writer, cards *cardEntries) {
	_, _ = fmt.Fprintln(writer, "File name:")
	scanner.Scan()
	fileName := scanner.Text()

	file, err := os.Create(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(writer, "%+v\n", err)
		return
	}
	defer func() { _ = file.Close() }()

	if err := json.NewEncoder(file).Encode(cards.cards); err != nil {
		_, _ = fmt.Fprintf(writer, "%+v\n", err)
		return
	}
	_, _ = fmt.Fprintf(writer, "%d cards have been saved.\n", len(cards.cards))
}

func addCard(scanner *bufio.Scanner, writer io.Writer, cards *cardEntries) {
	_, _ = fmt.Fprintln(writer, "The card:")
	var card string
	for scanner.Scan() {
		card = scanner.Text()
		if _, ok := cards.hasCard(card); ok {
			_, _ = fmt.Fprintf(writer, "The card \"%s\" already exists. Try again:\n", card)
			continue
		}
		break
	}

	_, _ = fmt.Fprintln(writer, "The definition of the card:")
	var definition string
	for scanner.Scan() {
		definition = scanner.Text()
		if _, ok := cards.hasDefinition(definition); ok {
			_, _ = fmt.Fprintf(writer, "The definition \"%s\" already exists. Try again:\n", definition)
			continue
		}
		break
	}

	cards.add(card, definition)
	_, _ = fmt.Fprintf(writer, "The pair (\"%s\":\"%s\") has been added.\n", card, definition)
}

func removeCard(scanner *bufio.Scanner, writer io.Writer, entries *cardEntries) {
	_, _ = fmt.Fprintln(writer, "Which card?")
	_ = scanner.Scan()
	card := scanner.Text()
	if _, ok := entries.hasCard(card); !ok {
		_, _ = fmt.Fprintf(writer, "Can't remove \"%s\": there is no such card.\n", card)
		return
	}

	entries.remove(card)
	_, _ = fmt.Fprintf(writer, "The card has been removed.\n")
}

func askCards(scanner *bufio.Scanner, writer io.Writer, entries *cardEntries) {
	_, _ = fmt.Fprintln(writer, "How many times to ask?")
	for scanner.Scan() {
		askCount, err := strconv.Atoi(scanner.Text())
		if err != nil {
			_, _ = fmt.Fprintln(writer, "Invalid input")
			continue
		}

		for i := 0; i < askCount; i++ {
			askCard(scanner, writer, entries, i%len(entries.cards))
		}
		break
	}
}

func askCard(scanner *bufio.Scanner, writer io.Writer, entries *cardEntries, i int) {
	_, _ = fmt.Fprintf(writer, "Print the definition of \"%s\":\n", entries.cards[i].Card)

	scanner.Scan()
	definition := scanner.Text()

	if definition == entries.cards[i].Definition {
		_, _ = fmt.Fprintln(writer, "Correct!")
		return
	}

	if card, ok := entries.hasDefinition(definition); ok {
		_, _ = fmt.Fprintf(writer, "Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", entries.cards[i].Definition, card)
		return
	}

	_, _ = fmt.Fprintf(writer, "Wrong. The right answer is \"%s\".\n", entries.cards[i].Definition)
}

func newCards() *cardEntries {
	return &cardEntries{cards: make([]CardEntry, 0)}
}

type cardEntries struct {
	cards []CardEntry
}

type CardEntry struct {
	Card       string `json:"card"`
	Definition string `json:"definition"`
}

func (ce *cardEntries) add(card, definition string) {
	ce.cards = append(ce.cards, CardEntry{Card: card, Definition: definition})
}

func (ce *cardEntries) hasCard(card string) (string, bool) {
	for _, c := range ce.cards {
		if c.Card == card {
			return c.Definition, true
		}
	}
	return "", false
}

func (ce *cardEntries) hasDefinition(definition string) (string, bool) {
	for _, c := range ce.cards {
		if c.Definition == definition {
			return c.Card, true
		}
	}
	return "", false
}

func (ce *cardEntries) upsertCard(entry CardEntry) {
	if _, ok := ce.hasCard(entry.Card); ok {
		ce.remove(entry.Card)
	}
	ce.add(entry.Card, entry.Definition)
}

func (ce *cardEntries) remove(card string) {
	for i, c := range ce.cards {
		if c.Card == card {
			ce.cards = append(ce.cards[:i], ce.cards[i+1:]...)
			return
		}
	}
}
