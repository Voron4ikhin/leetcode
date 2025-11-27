package internal

import "fmt"

func Start3Sum15() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	pFix := 0
	for pFix < len(nums) {
		checkSlice := nums[pFix+1:]
		twoNumsList := twoSumEx(checkSlice, -nums[pFix])
		for _, numsList := range twoNumsList {
			toAppend := []int{nums[pFix]}
			for _, num := range numsList {
				toAppend = append(toAppend, checkSlice[num])
			}
			if len(toAppend) == 3 {
				result = append(result, toAppend)
			}

		}
		pFix++
	}
	return result
}

func twoSumEx(nums []int, target int) [][]int {
	result := make([][]int, 0, 2)
	mapValues := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index, ok := mapValues[target-nums[i]]
		if !ok {
			mapValues[nums[i]] = i
			continue
		}
		toAppend := []int{index, i}
		result = append(result, toAppend)
	}
	return result
}
