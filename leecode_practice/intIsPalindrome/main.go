package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome3(21))
}

func isPalindrome(x int) bool {
	y := x
	n := 0
	for y > 0 {
		m := y % 10
		y = y / 10
		n = n*10 + m
	}
	if n == x {
		return true
	}

	return false
}

func isPalindrome2(x int) bool {
	sx := strconv.Itoa(x)
	l := len(sx)
	for i, _ := range sx[:l/2] {
		if sx[i] != sx[l-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome3(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	n := 0
	for x > n {
		n = n*10 + x%10
		x /= 10
	}
	return x==n || x==n/10
}
