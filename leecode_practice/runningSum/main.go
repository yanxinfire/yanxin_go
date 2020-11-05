package main

import "fmt"

func main() {
	a := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum2(a))
}

func runningSum(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	for i := l - 1; i >= 0; i-- {
		res[i] = sum
		sum = sum - nums[i]
	}
	return res
}

func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
