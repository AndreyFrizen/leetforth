package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v\n", intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Printf("%v\n", moveZeroes([]int{0, 1, 0, 3, 12}))
	fmt.Printf("%v\n", kElementMax([]int{1, 2, 3, 4, 5}, 3))
}

// twoSum finds two numbers in nums that add up to target and returns their indices.
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

// intersect finds the intersection of two sorted arrays and returns the common elements.
func intersect(nums1 []int, nums2 []int) []int {
	p1 := 0
	p2 := 0
	result := []int{}

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}
	}
	return result
}

// moveZeroes moves all zeroes in nums to the end while maintaining the relative order of non-zero elements.
func moveZeroes(nums []int) []int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}

	return nums
}

// kElMax finds the kth largest element in nums.
func kElementMax(nums []int, k int) int {
	windowSum := 0
	for i := range k {
		windowSum += nums[i]
	}
	currSum := windowSum
	for r := k; r < len(nums); r++ {
		l := r - k
		windowSum = windowSum + nums[r] - nums[l]
		if windowSum > currSum {
			currSum = windowSum
		}
	}
	return currSum
}
