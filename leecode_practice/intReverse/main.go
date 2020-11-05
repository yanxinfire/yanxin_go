package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 2147483647
	fmt.Println(reverse2(a))
}

func reverse(x int) int {
	b := x
	if x < 0 {
		b = 0 - x
	}
	s := strconv.FormatInt(int64(b), 10)
	var rs string
	for i := len(s) - 1; i >= 0; i-- {
		rs += fmt.Sprintf("%c", s[i])
	}
	ra, _ := strconv.ParseInt(rs, 10, 64)
	if x < 0 {
		ra = 0 - ra
	}
	if ra >= -2147483648 && ra <= 2147483647 {
		return int(ra)
	}
	return 0
}

func reverse2(x int) int {
	rev := 0
	for ; x != 0; {
		pop := x % 10
		x /= 10
		if rev < -2147483648/10 || rev > 2147483647/10 || (rev == -2147483648/10 && pop < -8) || (rev == 2147483647/10 && pop > 7) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
