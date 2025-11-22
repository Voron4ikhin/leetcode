package internal

import "fmt"

func StartLongestSubstringWithoutRepeatingCharacters3() {
	s := " "
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for index, _ := range s {
		current := 0
		checkMap := make(map[byte]int)
		for i := index; i < len(s); i++ {
			_, ok := checkMap[s[i]]
			if !ok {
				checkMap[s[i]] = 1
				current += 1
				continue
			}
			if current > res {
				res = current
			}
			break
		}
		if current > res {
			res = current
		}
	}
	return res
}
