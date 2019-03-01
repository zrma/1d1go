package strings

import (
	"bufio"
	"encoding/csv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/palindrome-index/problem", func() {
	It("문제를 풀었다", func() {
		actual := palindromeIndex("aaab")
		Expect(actual).Should(BeNumerically("==", 3))
		actual = palindromeIndex("baa")
		Expect(actual).Should(BeNumerically("==", 0))
		actual = palindromeIndex("aaa")
		Expect(actual).Should(BeNumerically("==", -1))
		actual = palindromeIndex("a")
		Expect(actual).Should(BeNumerically("==", -1))
		actual = palindromeIndex("abcadcdddaba")
		Expect(actual).Should(BeNumerically("==", -1))

		actual = palindromeIndex("quyjjdcgsvvsgcdjjyq")
		Expect(actual).Should(BeNumerically("==", 1))
		actual = palindromeIndex("hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh")
		Expect(actual).Should(BeNumerically("==", 8))
		actual = palindromeIndex("fgnfnidynhxebxxxfmxixhsruldhsaobhlcggchboashdlurshxixmfxxxbexhnydinfngf")
		Expect(actual).Should(BeNumerically("==", 33))
		actual = palindromeIndex("bsyhvwfuesumsehmytqioswvpcbxyolapfywdxeacyuruybhbwxjmrrmjxwbhbyuruycaexdwyfpaloyxbcpwsoiqtymhesmuseufwvhysb")
		Expect(actual).Should(BeNumerically("==", 23))
		actual = palindromeIndex("fvyqxqxynewuebtcuqdwyetyqqisappmunmnldmkttkmdlnmnumppasiqyteywdquctbeuwenyxqxqyvf")
		Expect(actual).Should(BeNumerically("==", 25))
		actual = palindromeIndex("mmbiefhflbeckaecprwfgmqlydfroxrblulpasumubqhhbvlqpixvvxipqlvbhqbumusaplulbrxorfdylqmgfwrpceakceblfhfeibmm")
		Expect(actual).Should(BeNumerically("==", 44))
		actual = palindromeIndex("tpqknkmbgasitnwqrqasvolmevkasccsakvemlosaqrqwntisagbmknkqpt")
		Expect(actual).Should(BeNumerically("==", 20))
		actual = palindromeIndex("lhrxvssvxrhl")
		Expect(actual).Should(BeNumerically("==", -1))
		actual = palindromeIndex("prcoitfiptvcxrvoalqmfpnqyhrubxspplrftomfehbbhefmotfrlppsxburhyqnpfmqlaorxcvtpiftiocrp")
		Expect(actual).Should(BeNumerically("==", 14))
		actual = palindromeIndex("kjowoemiduaaxasnqghxbxkiccikxbxhgqnsaxaaudimeowojk")
		Expect(actual).Should(BeNumerically("==", -1))
	})

	Context("equalPrefix 함수는", func() {
		Context("1글자 문자열 두 개가", func() {
			It("같으면 true를 반환한다.", func() {
				actual := equalPrefix("a", "a")
				Expect(actual).Should(BeTrue())
			})

			It("다르면 false를 반환한다.", func() {
				actual := equalPrefix("a", "b")
				Expect(actual).Should(BeFalse())
			})
		})

		Context("2글자 이상의 길이의 문자열 두 개의 앞 2글자", func() {
			It("같으면 true를 반환한다.", func() {
				actual := equalPrefix("abcd", "abab")
				Expect(actual).Should(BeTrue())
			})

			It("다르면 false를 반환한다.", func() {
				actual := equalPrefix("abc", "bbc")
				Expect(actual).Should(BeFalse())
			})
		})
	})

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long string", func() {
			file, err := os.Open("./test_data/palindrome_index.csv")
			Expect(err).Should(BeNil())

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			Expect(err).Should(BeNil())

			var arr []string
			for _, row := range rows {
				arr = append(arr, row[0])
			}

			expected := []int32{
				16722,
				76661,
				10035,
				57089,
				46674,
			}

			Expect(len(arr)).Should(Equal(5))
			for i, s := range arr {
				actual := palindromeIndex(s)
				Expect(actual).Should(BeNumerically("==", expected[i]))
			}
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 5), "시간 초과")
	}, 10)
})
