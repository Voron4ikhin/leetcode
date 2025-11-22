package internal

import "fmt"

func StartLongestPalindromicSubstring5() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
	findPalindrome := func(left, right int) {
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			if len(s[left:right+1]) > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		findPalindrome(i, i)
		findPalindrome(i, i+1)
	}
	return res
}
