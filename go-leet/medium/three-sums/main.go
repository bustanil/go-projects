package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	for i := 0; i < len(result); i++ {
		fmt.Printf("[%d, %d, %d] = %d\n", result[i][0], result[i][1], result[i][2], result[i][0]+result[i][1]+result[i][2])
	}
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var tries = make(map[int]bool)
	var set = make(map[string][]int)
	for i := 0; i < len(nums)-2; i++ {
		if tries[nums[i]] {
			continue
		}
		tries[nums[i]] = true
		for j := i + 1; j < len(nums)-1; j++ {
			target := 0 - (nums[i] + nums[j])
			haystack := nums[j+1:]
			n, found := slices.BinarySearch(haystack, target)
			if found {
				triplet := []int{nums[i], nums[j], haystack[n]}
				set[fmt.Sprintf("%d%d%d", triplet[0], triplet[1], triplet[2])] = triplet
			}
		}
	}
	result := make([][]int, 0)
	for _, values := range set {
		result = append(result, values)
	}
	return result
}
