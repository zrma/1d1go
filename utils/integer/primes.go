package integer

func BuildSieveOfEratosthenes(maxLen int) []bool {
	sieveOfEratosthenes := make([]bool, maxLen)
	for i := 2; i*i < maxLen; i++ {
		if !sieveOfEratosthenes[i] {
			for j := i * i; j < maxLen; j += i {
				sieveOfEratosthenes[j] = true
			}
		}
	}
	sieveOfEratosthenes[0] = true
	sieveOfEratosthenes[1] = true
	return sieveOfEratosthenes
}
