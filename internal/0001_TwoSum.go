package internal

import "fmt"

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func StartTwoSum1() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	mapValues := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index, ok := mapValues[target-nums[i]]
		if !ok {
			mapValues[nums[i]] = i
			continue
		}
		result = append(result, index, i)
	}
	return result
}
