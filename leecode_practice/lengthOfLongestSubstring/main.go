package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring2(s))
}

func lengthOfLongestSubstring(s string) int {
	cnt := 0
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
	return ans
}

func lengthOfLongestSubstring2(s string) int {
	sublen, idx := 0, 0
	n := len(s)
	for i := 0; i < n; {
		if i==n-1{
			return max(sublen,1)
		}
		for end := i + 1; end < n; end++ {
			if idx = strings.Index(s[i:end], s[end:end+1]); idx >= 0 {
				sublen = max(end-i, sublen)
				i += idx + 1
				break
			}
			if end == n-1 {
				return max(end-i+1, sublen)
			}
		}
	}
	return sublen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
