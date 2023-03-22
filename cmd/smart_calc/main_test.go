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
Invalid expression
Invalid expression
15
Invalid expression
-22
Invalid expression
5
Unknown command
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
