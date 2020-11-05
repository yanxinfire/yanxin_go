package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 2, 4, 10, 6, 10, -8, 8}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	sntc := make([]int, len(nums))
	ordernums := make([]int, len(nums))
	copy(ordernums, nums)
	//var temp int
	//for m := 0; m < len(ordernums)-1; m++ {
	//	for n := 0; n < len(ordernums)-m-1; n++ {
	//		if ordernums[n] > ordernums[n+1] {
	//			temp = ordernums[n]
	//			ordernums[n] = ordernums[n+1]
	//			ordernums[n+1] = temp
	//		}
	//	}
	//}
	sort.Slice(ordernums, func(i, j int) bool {
		return ordernums[i] < ordernums[j]
	})
	hashmap := make(map[int]int)
	for i, v := range ordernums {
		if _, ok := hashmap[v]; !ok {
			hashmap[v] = i
		}
	}
	for i, v := range nums {
		sntc[i] = hashmap[v]
	}
	return sntc
}

func smallerNumbersThanCurrent2(nums []int) []int {
	sntc := make([]int, len(nums))
	freq := [101]int{}
	for _, v := range nums {
		freq[v]++
	}
	for i := 0; i < 100; i++ {
		freq[i+1] += freq[i]
	}
	for i, v := range nums {
		if v > 0 {
			sntc[i] = freq[v-1]
		}
	}
	return sntc
}
