package internal

import (
	"fmt"
	"math"
)

func StartReverseInteger7() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	positive := true
	if x < 0 {
		positive = false
		x = -x
	}
	var result int
	for x > 0 {
		mod := x % 10
		result = result*10 + mod
		if result > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	if !positive {
		return -result
	}
	return result
}
