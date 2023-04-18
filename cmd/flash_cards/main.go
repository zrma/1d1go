package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	importFile := flag.String("import_from", "", "import file")
	exportFile := flag.String("export_to", "", "export file")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	FlashCards(scanner, os.Stdout, *importFile, *exportFile)
}

func newReadWriter(scanner *bufio.Scanner, writer io.Writer) *readWriter {
	return &readWriter{
		dstWriter:  new(strings.Builder),
		srcScanner: scanner,
		srcWriter:  writer,
	}
}

type readWriter struct {
	dstWriter  *strings.Builder
	srcScanner *bufio.Scanner
	srcWriter  io.Writer
}

func (l *readWriter) Write(p []byte) (n int, err error) {
	if _, err := l.dstWriter.Write(p); err != nil {
		return 0, err
	}
	return l.srcWriter.Write(p)
}

func (l *readWriter) Scan() bool {
	return l.srcScanner.Scan()
}

func (l *readWriter) Text() string {
	s := l.srcScanner.Text()
	_, _ = fmt.Fprintln(l.dstWriter, s)
	return s
}

func FlashCards(scanner *bufio.Scanner, writer io.Writer, importFile string, exportFile string) {
	cards := newCards()

	rw := newReadWriter(scanner, writer)

	if importFile != "" {
		importCardsFromFile(importFile, rw, cards)
	}

	for {
		_, _ = fmt.Fprintln(rw, "Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")

		if !rw.Scan() {
			return
		}
		action := rw.Text()

		switch action {
		case "add":
			addCard(rw, cards)
		case "remove":
			removeCard(rw, cards)
		case "import":
			importCards(rw, cards)
		case "export":
			exportCards(rw, cards)
		case "ask":
			askCards(rw, cards)
		case "exit":
			_, _ = fmt.Fprintln(rw, "Bye bye!")

			if exportFile != "" {
				exportCardsToFile(exportFile, rw, cards)
			}
			return
		case "log":
			logging(rw)
		case "hardest card":
			hardestCard(rw, cards)
		case "reset stats":
			resetStats(rw, cards)
		}

		_, _ = fmt.Fprintln(rw)
	}
}

func importCards(rw *readWriter, cards *cardEntries) {
	_, _ = fmt.Fprintln(rw, "File name:")
	rw.Scan()
	fileName := rw.Text()

	importCardsFromFile(fileName, rw, cards)
}

func importCardsFromFile(fileName string, rw *readWriter, cards *cardEntries) {
	file, err := os.Open(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(rw, "File not found.\n")
		return
	}
	defer func() { _ = file.Close() }()

	var newCards []CardEntry
	if err := json.NewDecoder(file).Decode(&newCards); err != nil {
		_, _ = fmt.Fprintf(rw, "%+v\n", err)
		return
	}

	for _, card := range newCards {
		cards.upsertCard(card)
	}
	_, _ = fmt.Fprintf(rw, "%d cards have been loaded.\n", len(newCards))
}

func exportCards(rw *readWriter, cards *cardEntries) {
	_, _ = fmt.Fprintln(rw, "File name:")
	rw.Scan()
	fileName := rw.Text()

	exportCardsToFile(fileName, rw, cards)
}

func exportCardsToFile(fileName string, rw *readWriter, cards *cardEntries) {
	file, err := os.Create(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(rw, "%+v\n", err)
		return
	}
	defer func() { _ = file.Close() }()

	if err := json.NewEncoder(file).Encode(cards.cards); err != nil {
		_, _ = fmt.Fprintf(rw, "%+v\n", err)
		return
	}
	_, _ = fmt.Fprintf(rw, "%d cards have been saved.\n", len(cards.cards))
}

func addCard(rw *readWriter, cards *cardEntries) {
	_, _ = fmt.Fprintln(rw, "The card:")
	var card string
	for rw.Scan() {
		card = rw.Text()
		if _, ok := cards.hasCard(card); ok {
			_, _ = fmt.Fprintf(rw, "The card \"%s\" already exists. Try again:\n", card)
			continue
		}
		break
	}

	_, _ = fmt.Fprintln(rw, "The definition of the card:")
	var definition string
	for rw.Scan() {
		definition = rw.Text()
		if _, ok := cards.hasDefinition(definition); ok {
			_, _ = fmt.Fprintf(rw, "The definition \"%s\" already exists. Try again:\n", definition)
			continue
		}
		break
	}

	cards.add(card, definition, 0)
	_, _ = fmt.Fprintf(rw, "The pair (\"%s\":\"%s\") has been added.\n", card, definition)
}

func removeCard(rw *readWriter, entries *cardEntries) {
	_, _ = fmt.Fprintln(rw, "Which card?")
	_ = rw.Scan()
	card := rw.Text()
	if _, ok := entries.hasCard(card); !ok {
		_, _ = fmt.Fprintf(rw, "Can't remove \"%s\": there is no such card.\n", card)
		return
	}

	entries.remove(card)
	_, _ = fmt.Fprintf(rw, "The card has been removed.\n")
}

func askCards(rw *readWriter, entries *cardEntries) {
	_, _ = fmt.Fprintln(rw, "How many times to ask?")
	for rw.Scan() {
		askCount, err := strconv.Atoi(rw.Text())
		if err != nil {
			_, _ = fmt.Fprintln(rw, "Invalid input")
			continue
		}

		for i := 0; i < askCount; i++ {
			askCard(rw, entries, i%len(entries.cards))
		}
		break
	}
}

func askCard(rw *readWriter, entries *cardEntries, i int) {
	_, _ = fmt.Fprintf(rw, "Print the definition of \"%s\":\n", entries.cards[i].Card)

	rw.Scan()
	definition := rw.Text()

	if definition == entries.cards[i].Definition {
		_, _ = fmt.Fprintln(rw, "Correct!")
		return
	}

	entries.cards[i].ErrorCount++

	if card, ok := entries.hasDefinition(definition); ok {
		_, _ = fmt.Fprintf(rw, "Wrong. The right answer is \"%s\", but your definition is correct for \"%s\" card.\n", entries.cards[i].Definition, card)
		return
	}

	_, _ = fmt.Fprintf(rw, "Wrong. The right answer is \"%s\".\n", entries.cards[i].Definition)
}

func hardestCard(rw *readWriter, cards *cardEntries) {
	if len(cards.cards) == 0 {
		_, _ = fmt.Fprintln(rw, "There are no cards with errors.")
		return
	}

	maxErrorCount := 0
	for _, card := range cards.cards {
		if card.ErrorCount > maxErrorCount {
			maxErrorCount = card.ErrorCount
		}
	}

	if maxErrorCount == 0 {
		_, _ = fmt.Fprintln(rw, "There are no cards with errors.")
		return
	}

	var hardestCards []string
	for _, card := range cards.cards {
		if card.ErrorCount == maxErrorCount {
			hardestCards = append(hardestCards, card.Card)
		}
	}

	if len(hardestCards) == 1 {
		_, _ = fmt.Fprintf(rw, "The hardest card is \"%s\". You have %d errors answering it.\n", hardestCards[0], maxErrorCount)
		return
	}

	_, _ = fmt.Fprintf(rw, "The hardest cards are \"%s\". You have %d errors answering them.\n", strings.Join(hardestCards, "\", \""), maxErrorCount)
}

func resetStats(writer io.Writer, cards *cardEntries) {
	for i := range cards.cards {
		cards.cards[i].ErrorCount = 0
	}
	_, _ = fmt.Fprintln(writer, "Card statistics have been reset.")
}

func logging(rw *readWriter) {
	_, _ = fmt.Fprintln(rw, "File name:")
	_ = rw.Scan()
	fileName := rw.Text()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		_, _ = fmt.Fprintf(rw, "%+v\n", err)
		return
	}
	defer func() { _ = file.Close() }()

	_, _ = file.WriteString(rw.dstWriter.String())
	rw.dstWriter.Reset()

	_, _ = fmt.Fprintln(rw, "The log has been saved.")
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
	ErrorCount int    `json:"error_count"`
}

func (ce *cardEntries) add(card, definition string, errorCount int) {
	ce.cards = append(ce.cards, CardEntry{Card: card, Definition: definition, ErrorCount: errorCount})
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
	ce.add(entry.Card, entry.Definition, entry.ErrorCount)
}

func (ce *cardEntries) remove(card string) {
	for i, c := range ce.cards {
		if c.Card == card {
			ce.cards = append(ce.cards[:i], ce.cards[i+1:]...)
			return
		}
	}
}
