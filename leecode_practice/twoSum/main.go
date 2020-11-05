package main

import "fmt"

func main() {
	t := []int{-3, 3, 4, 3, 90}
	r := twoSum3(t, 6)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	newnums := [][]int{}
	results := []int{}
	var tmp []int

	for i, v := range nums {
		newnums = append(newnums, []int{i, v})
	}

	for m := 0; m < len(newnums)-1; m++ {
		for n := 0; n < len(newnums)-1-m; n++ {
			if newnums[n][1] > newnums[n+1][1] {
				tmp = newnums[n]
				newnums[n] = newnums[n+1]
				newnums[n+1] = tmp
			}
		}
	}

	for x := 0; x < len(newnums)-1; x++ {
		for y := x + 1; y < len(newnums); y++ {
			if newnums[x][1]+newnums[y][1] == target {
				results = append(results, newnums[x][0])
				results = append(results, newnums[y][0])
				return results
			}
		}
	}

	return results
}

func twoSum2(nums []int, target int) []int {
	numsmap := make(map[int][]int)
	for i, v := range nums {
		numsmap[v] = append(numsmap[v], i)
	}
	for k, v := range numsmap {
		if v1, ok := numsmap[target-k]; ok {
			if target-k == k && len(numsmap[k]) >= 2 {
				return []int{v1[0], v1[1]}
			} else if target-k != k {
				return []int{v[0], v1[0]}
			}
		}
	}
	return []int{}
}

func twoSum3(nums []int, target int) []int {
	numsmap := make(map[int]int)
	for i, v := range nums {
		if p, ok := numsmap[target-v]; ok {
			return []int{p, i}
		}
		numsmap[v] = i
	}
	return nil
}
