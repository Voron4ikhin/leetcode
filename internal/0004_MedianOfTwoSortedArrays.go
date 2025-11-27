package internal

import "fmt"

func StartMedianOfTwoSortedArrays4() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)
	sumLen := lenNums1 + lenNums2
	p1 := 0
	p2 := 0
	canAdd1 := true
	canAdd2 := true
	isOdd := false
	if sumLen%2 == 1 {
		isOdd = true
	}
	for p1+p2 < sumLen/2-1 {
		if p1 == len(nums1)-1 {
			canAdd1 = false
		}
		if p2 == len(nums2)-1 {
			canAdd2 = false
		}
		if nums1[p1] < nums2[p2] {
			if canAdd1 {
				p1++
				continue
			}
			p2++
			continue
		}
		if canAdd2 {
			p2++
			continue
		}
		p1++
		if p1 == len(nums1)-1 {
			canAdd1 = false
		}
		if p2 == len(nums2)-1 {
			canAdd2 = false
		}
		continue
	}

	if isOdd {
		fmt.Println("тут")
		fmt.Println(p1)
		fmt.Println(p2)
		if canAdd1 {
			return float64(nums1[p1])
		}
		if canAdd2 {
			return float64(nums2[p2])
		}
	}

	if canAdd1 && canAdd2 {
		return float64(nums1[p1]+nums2[p2]) / 2
	}

	if canAdd1 {
		if p1 == 0 {
			return float64(nums2[p2]+nums1[p1]) / 2
		}
		return float64(nums1[p1]+nums1[p1+1]) / 2
	}

	if canAdd2 {
		if p2 == 0 {
			return float64(nums1[p1]+nums2[p2]) / 2
		}
		return float64(nums2[p2]+nums2[p2+1]) / 2
	}

	return 0
}
