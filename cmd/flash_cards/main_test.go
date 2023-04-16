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
			`add
France
Paris
add
France
Great Britain
Paris
London
remove
France
remove
Wakanda
exit`,
			`Input the action (add, remove, import, export, ask, exit):
The card:
The definition of the card:
The pair ("France":"Paris") has been added.

Input the action (add, remove, import, export, ask, exit):
The card:
The card "France" already exists. Try again:
The definition of the card:
The definition "Paris" already exists. Try again:
The pair ("Great Britain":"London") has been added.

Input the action (add, remove, import, export, ask, exit):
Which card?
The card has been removed.

Input the action (add, remove, import, export, ask, exit):
Which card?
Can't remove "Wakanda": there is no such card.

Input the action (add, remove, import, export, ask, exit):
Bye bye!
`,
		},
		{
			`import
ghost_file.txt
add
Japan
Tokyo
add
Russia
UpdateMeFromFile
import
capitals.txt
ask
2
Tokyo
Moscow
export
capitalsNew.txt
exit`,
			`Input the action (add, remove, import, export, ask, exit):
File name:
File not found.

Input the action (add, remove, import, export, ask, exit):
The card:
The definition of the card:
The pair ("Japan":"Tokyo") has been added.

Input the action (add, remove, import, export, ask, exit):
The card:
The definition of the card:
The pair ("Russia":"UpdateMeFromFile") has been added.

Input the action (add, remove, import, export, ask, exit):
File name:
28 cards have been loaded.

Input the action (add, remove, import, export, ask, exit):
How many times to ask?
Print the definition of "Japan":
Correct!
Print the definition of "Russia":
Correct!

Input the action (add, remove, import, export, ask, exit):
File name:
29 cards have been saved.

Input the action (add, remove, import, export, ask, exit):
Bye bye!
`,
		},
		{
			`add
a brother of one's parent
uncle
add
a part of the body where the foot and the leg meet
ankle
ask
6
ankle
??
uncle
ankle
??
uncle
exit`,
			`Input the action (add, remove, import, export, ask, exit):
The card:
The definition of the card:
The pair ("a brother of one's parent":"uncle") has been added.

Input the action (add, remove, import, export, ask, exit):
The card:
The definition of the card:
The pair ("a part of the body where the foot and the leg meet":"ankle") has been added.

Input the action (add, remove, import, export, ask, exit):
How many times to ask?
Print the definition of "a brother of one's parent":
Wrong. The right answer is "uncle", but your definition is correct for "a part of the body where the foot and the leg meet".
Print the definition of "a part of the body where the foot and the leg meet":
Wrong. The right answer is "ankle".
Print the definition of "a brother of one's parent":
Correct!
Print the definition of "a part of the body where the foot and the leg meet":
Correct!
Print the definition of "a brother of one's parent":
Wrong. The right answer is "uncle".
Print the definition of "a part of the body where the foot and the leg meet":
Wrong. The right answer is "ankle", but your definition is correct for "a brother of one's parent".

Input the action (add, remove, import, export, ask, exit):
Bye bye!
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
