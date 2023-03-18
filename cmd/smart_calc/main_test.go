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
			`8
-2 + 4 - 5 + 6
9 +++ 10 -- 8
3 --- 5
14       -   12
/exit`,
			`8
3
27
-2
2
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
