package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	for l < r {
		sum := nums[l] + nums[r]

		if sum == target {
			return []int{l, r}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
