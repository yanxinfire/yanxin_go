package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcde"))
}

func lengthOfLongestSubstring(s string) int {
	cnt:=0
	hashmap := make(map[byte]int)
	ans, rk := 0, -1
	n := len(s)
	for i := 0; i < n; i++ {
		cnt++
		if i != 0 {
			delete(hashmap, s[i-1])
		}
		for rk+1 < n && hashmap[s[rk+1]] == 0 {
			hashmap[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
		if rk == (n - 1) {
			break
		}
	}
	fmt.Println(cnt)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
