package p4400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve4458(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4458")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`5
powdered Toast Man
skeletor
Electra Woman and Dyna Girl
she-Ra Princess of Power
darth Vader`,
			`Powdered Toast Man
Skeletor
Electra Woman and Dyna Girl
She-Ra Princess of Power
Darth Vader
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve4458(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
