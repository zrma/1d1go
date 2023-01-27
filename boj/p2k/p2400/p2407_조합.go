package p2400

import (
	"fmt"
	"io"
	"math/big"
)

func Solve2407(reader io.Reader, writer io.Writer) {
	var n, m big.Int
	_, _ = fmt.Fscan(reader, &n, &m)

	const maxLen = 1000 + 1
	cache := make([][]*big.Int, maxLen)
	for i := range cache {
		cache[i] = make([]*big.Int, maxLen)
	}

	res := combination(&n, &m, cache)
	_, _ = fmt.Fprint(writer, res.String())
}

func combination(n *big.Int, m *big.Int, cache [][]*big.Int) *big.Int {
	if n.Cmp(m) < 0 {
		return big.NewInt(0)
	}
	if n.Cmp(m) == 0 {
		return big.NewInt(1)
	}
	if m.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	if m.Cmp(big.NewInt(1)) == 0 {
		return n
	}

	if cache[n.Int64()][m.Int64()] != nil {
		return cache[n.Int64()][m.Int64()]
	}

	res := big.NewInt(0)
	res = res.Add(res, combination(new(big.Int).Sub(n, big.NewInt(1)), new(big.Int).Sub(m, big.NewInt(1)), cache))
	res = res.Add(res, combination(new(big.Int).Sub(n, big.NewInt(1)), m, cache))
	cache[n.Int64()][m.Int64()] = res

	return res
}
