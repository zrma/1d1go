package hacker_rank_all_womens_codesprint_2019

func alphaBeta(pile []int32) int32 {
	var best, better int32
	for _, p := range pile {
		if p > best {
			best, better = p, best
		} else if p < best && p > better {
			better = p
		}
	}
	return better
}
