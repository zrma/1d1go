package strings

import (
	"bytes"
)

func stringConstruction(s string) int32 {
	var p []byte
	var cost int32
	for i := 0; i < len(s); i++ {
		c := s[i]

		if !bytes.ContainsRune(p, rune(c)) {
			cost++
		}
		p = append(p, c)
	}

	return cost
}
