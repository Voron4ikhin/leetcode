package internal

import "fmt"

func StartRomanToInteger13() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	p1 := 'M'
	result := 0
	for _, v := range s {
		value := romanMap[v]
		if value > romanMap[p1] {
			result -= 2 * romanMap[p1]
			result += value
			continue
		}
		result += value
		p1 = v
	}
	return result
}
