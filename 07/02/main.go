package main

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(arr, 3)
}

func rotate(nums []int, k int) {
	lNums := len(nums)
	if k > lNums {
		n := k / lNums
		k -= lNums * n
	}
	if lNums == k || k == 0 {
		return
	}
	buf := make([]int, k)

	j := 0
	for i := lNums - k; i < lNums; i++ {
		buf[j] = nums[i]
		j++
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = nums[i-k]
		if i == k {
			break
		}
	}
	copy(nums,buf)
}

