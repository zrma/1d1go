package p3000

import (
	"fmt"
	"io"
)

func Solve3003(reader io.Reader, writer io.Writer) {
	var king, queen, rook, bishop, knight, pawn int
	_, _ = fmt.Fscan(reader, &king, &queen, &rook, &bishop, &knight, &pawn)

	const (
		wantKing   = 1
		wantQueen  = 1
		wantRook   = 2
		wantBishop = 2
		wantKnight = 2
		wantPawn   = 8
	)

	_, _ = fmt.Fprint(writer,
		wantKing-king,
		wantQueen-queen,
		wantRook-rook,
		wantBishop-bishop,
		wantKnight-knight,
		wantPawn-pawn,
	)
}
