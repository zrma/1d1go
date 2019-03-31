package strings

// 주어진 k번의 기회 안에서 회문으로 만들기
func makeSimplePalindrome(s string, n, k int32) ([]byte, int32, bool) {
	var i1, i2 int32
	result := []byte(s)
	for i1, i2 = 0, n-1; i1 < i2; i1, i2 = i1+1, i2-1 {
		if result[i1] > result[i2] {
			result[i2] = result[i1]
			k--

			if k < 0 {
				return nil, -1, false
			}
		} else if result[i2] > result[i1] {
			result[i1] = result[i2]
			k--

			if k < 0 {
				return nil, -1, false
			}
		}
	}
	return result, k, true
}

// 보너스 기회로 더 큰 수로 만들기
func makeHighestValuePalindrome(s string, b []byte, n, k int32) []byte {
	var i1, i2 int32
	for i1, i2 = 0, n-1; i1 <= i2; i1, i2 = i1+1, i2-1 {
		if i1 == i2 && k > 0 {
			b[i1] = '9'
		}

		// 만약 9라면 이미 앞쪽에서 양쪽 모두 9 회문으로 바뀐 것이므로 패스
		if b[i1] != '9' {
			if (b[i1] == s[i1] && b[i2] == s[i2]) && k >= 2 {
				// 두 수 모두 변화가 없었다면,
				// k에서 2번의 기회를 사용해 양쪽 모두 9로 바꾸자.
				b[i1] = '9'
				b[i2] = '9'
				k -= 2
			} else if (b[i1] != s[i1] || b[i2] != s[i2]) && k >= 1 {
				// 둘 중 한 수가 이미 바뀐 경우
				// 회문 만들기 단게에서 9가 아닌 다른 수로 바꾸었을텐데,
				// 그 때 '9'로 바꿨다고 가정하고,
				// k에서 1번의 기회를 사용해 나머지 한 수도 9로 바꾸자.
				b[i1] = '9'
				b[i2] = '9'
				k -= 1
			}
		}
	}
	return b
}

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

	b, k, ok := makeSimplePalindrome(s, n, k)
	if !ok {
		return "-1"
	}

	b = makeHighestValuePalindrome(s, b, n, k)
	return string(b)
}
