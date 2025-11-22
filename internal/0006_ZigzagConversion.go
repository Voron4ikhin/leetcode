package internal

import "fmt"

func StartZigzagConversion6() {
	s := "ABC"
	numRows := 2
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	stringSlice := make([]string, numRows)
	ptr, reverse := 0, false

	for idx := 0; idx < len(s); idx++ {
		ss := string(s[idx])
		stringSlice[ptr] += ss

		if ptr == numRows-1 {
			reverse = true
			ptr -= 1
			continue
		}
		if !reverse {
			ptr += 1
		}
		if ptr == 0 {
			reverse = false
			ptr += 1
			continue
		}
		if reverse {
			ptr -= 1
		}

	}

	finalStr := ""
	for _, v := range stringSlice {
		finalStr += v
	}
	stringSlice = nil
	return finalStr
}
