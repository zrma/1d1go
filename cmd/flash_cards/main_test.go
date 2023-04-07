package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	t.Log("https://hyperskill.org/projects/224")

	tests := []struct {
		give string
		want string
	}{
		{
			`print()
outputs text
outputs text`,
			`Your answer is right!
`,
		},
		{
			`Jetbrains
A place for people who love to code
A place for people who hate to code`,
			`Your answer is wrong...
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			FlashCards(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
