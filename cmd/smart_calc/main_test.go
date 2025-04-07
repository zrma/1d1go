package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	t.Log("https://hyperskill.org/projects/278")

	tests := []struct {
		give string
		want string
	}{
		{
			`8 + 7 - 4
abc
123+
+15
18 22

-22
22-
8 --- 3
/go
/exit`,
			`11
unknown variable
invalid expression
15
invalid expression
-22
invalid expression
5
unknown command
Bye!
`,
		},
		{
			`n = 3
m=4
a  =   5
b = a
v=   7
n =9
count = 10
a = 1
a = 2
a = 3
a
X = 10
X
Y= 5
Y
z = X
Z = Y
z
Z`,
			`3
10
5
10
5
`,
		},
		{
			`a2a
n22`,
			`invalid identifier
invalid identifier
`,
		},
		{
			`a = 8
b = c
e`,
			`unknown variable
unknown variable
`,
		},
		{
			`a1 = 8
n1 = a2a
n = a2a
a = 7 = 8`,
			`invalid identifier
invalid identifier
invalid assignment
invalid assignment
`,
		},
		{
			`2^2
2*2^3`,
			`4
16
`,
		},
		{
			`8 * 3 + 12 * (4 - 2)
2 - 2 + 3
4 * (2 + 3
-10
a=4
b=5
c=6
a*2+b*3+c*(2+3)
1 +++ 2 * 3 -- 4
3 *** 5
4+3)
/command
/exit`,
			`48
3
invalid expression
-10
53
11
invalid expression
invalid expression
unknown command
Bye!
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			SmartCalc(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
