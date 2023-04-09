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
			`2
print()
outputs text
print()
str()
outputs text
converts to a string
outputs text
converts to a string`,
			`Input the number of cards:
The term for card #1:
The definition for card #1:
The term for card #2:
The term "print()" already exists. Try again:
The definition for card #2:
The definition "outputs text" already exists. Try again:
Print the definition of "print()":
Correct!
Print the definition of "str()":
Correct!
`,
		},
		{
			`2
uncle
a brother of one's parent
ankle
a part of the body where the foot and the leg meet
a part of the body where the foot and the leg meet
???`,
			`Input the number of cards:
The term for card #1:
The definition for card #1:
The term for card #2:
The definition for card #2:
Print the definition of "uncle":
Wrong. The right answer is "a brother of one's parent", but your definition is correct for "ankle".
Print the definition of "ankle":
Wrong. The right answer is "a part of the body where the foot and the leg meet".
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
