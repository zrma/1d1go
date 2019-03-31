package strings

func highestValuePalindrome(s string, n, k int32) string {
	if n == 0 {
		return "-1"
	}
	if n == 1 {
		if k <= 0 {
			return s
		}

		return "9"
	}

	var i1, i2 int32
	result := []byte(s)
	// 주어진 k번의 기회 안에서 회문으로 만들기
	for i1, i2 = 0, n-1; i1 < i2; i1, i2 = i1+1, i2-1 {
		if result[i1] > result[i2] {
			result[i2] = result[i1]
			k--

			if k < 0 {
				return "-1"
			}
		} else if result[i2] > result[i1] {
			result[i1] = result[i2]
			k--

			if k < 0 {
				return "-1"
			}
		}
	}

	i1, i2 = 0, n-1
	for i1, i2 = 0, n-1; i1 <= i2; i1, i2 = i1+1, i2-1 {
		if i1 == i2 && k > 0 {
			result[i1] = '9'
		}

		if result[i1] != '9' {
			if (result[i1] == s[i1] && result[i2] == s[i2]) && k >= 2 {
				// 두 수 모두 변화가 없었다면,
				//   k 중 2번의 기회를 소진 해 양쪽 모두 9로 바꾸자.
				result[i1] = '9'
				result[i2] = '9'
				k -= 2
			} else if (result[i1] != s[i1] || result[i2] != s[i2]) && k >= 1 {
				// 둘 중 한 수가 이미 바뀐 경우,
				//   (이미 회문이므로) 이전에 바꾼 수를 회문의 수가 아닌 '9'로 바꿨다고 가정하고,
				//   k 중 1번의 기회를 소진 해 나머지 한 수도 9로 바꾸자.
				result[i1] = '9'
				result[i2] = '9'
				k -= 1
			}
		}
	}

	return string(result)
}
