package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(findUnsortedSubarray(arr))
}

func findUnsortedSubarray(nums []int) int {
	lNums := len(nums)
	countLeft, countRight := 0, 0
	if lNums == 1 {
		return 0
	}
	sorted := mergeSort(nums)

	for i := 0; i < lNums; i++ {
		if nums[i] == sorted[i] {
			countLeft++
		} else {
			break
		}
	}
	if countLeft == lNums {
		return 0
	}
	for i := lNums - 1; i >= 0; i-- {
		if nums[i] == sorted[i] {
			countRight++
		} else {
			break
		}
	}
	return lNums - countLeft - countRight
}

func mergeSort(nums []int) []int {
	lNums := len(nums)

	if lNums == 1 {
		return nums
	}

	middle := lNums / 2

	left := make([]int, middle)
	right := make([]int, lNums-middle)

	for i := 0; i < lNums; i++ {
		if i < middle {
			left[i] = nums[i]
		} else {
			right[i-middle] = nums[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
