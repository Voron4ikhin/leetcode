package internal

import "fmt"

func StartLongestCommonPrefix14() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longestPrefix := strs[0]
	for _, v := range strs {
		p := 0
		for p < len(longestPrefix) && p < len(v) {
			if v[p] == longestPrefix[p] {
				p++
				continue
			}
			longestPrefix = v[:p]
			break
		}
		longestPrefix = v[:p]
	}
	return longestPrefix
}
