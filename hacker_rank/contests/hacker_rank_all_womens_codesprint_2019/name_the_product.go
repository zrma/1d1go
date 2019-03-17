package hacker_rank_all_womens_codesprint_2019

func productName(names []string) string {
	//noinspection SpellCheckingInspection
	chars := "abcdefghijklmnopqrstuvwxyz"

	size := len(names)
	first := names[0]
	var arr []byte
	for i := 0; i < len(first); i++ {
		var c byte
		var maxDiff int
		for j := 0; j < len(chars); j++ {
			var curDiff int
			for k := 0; k < size; k++ {
				if names[k][i] != chars[j] {
					curDiff++
				}
			}
			if maxDiff <= curDiff {
				maxDiff = curDiff
				c = chars[j]
			}
		}
		arr = append(arr, c)
	}

	return string(arr)
}
