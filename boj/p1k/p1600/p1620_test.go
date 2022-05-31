package p1600_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1600"
	"1d1go/utils"
)

func TestSolve1620(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1620")

	//goland:noinspection SpellCheckingInspection
	const (
		s = `26 5
Bulbasaur
Ivysaur
Venusaur
Charmander
Charmeleon
Charizard
Squirtle
Wartortle
Blastoise
Caterpie
Metapod
Butterfree
Weedle
Kakuna
Beedrill
Pidgey
Pidgeotto
Pidgeot
Rattata
Raticate
Spearow
Fearow
Ekans
Arbok
Pikachu
Raichu
25
Raichu
3
Pidgey
Kakuna`
		want = `Pikachu
26
Venusaur
16
14
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p1600.Solve1620(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
