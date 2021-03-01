package maxsatisfied

import "fmt"

func MaxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(grumpy)
	sum1 := 0
	incr := 0

	for i := 0; i < n; i++ {
		sum1 += customers[i] * (1 - grumpy[i])
	}

	for i := 0; i < X; i++ {
		incr += customers[i] * grumpy[i]
	}
	maxincr := incr
	for i := 1; i <= n-X; i++ {
		incr = incr - customers[i-1]*grumpy[i-1] + customers[i+X-1]*grumpy[i+X-1]
		if incr > maxincr {
			maxincr = incr
		}
	}
	return sum1 + maxincr
}

func AmaxSatisfied(customers []int, grumpy []int, X int) int {
	total := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := X; i < n; i++ {
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	fmt.Println(total + maxIncrease)
	return total + maxIncrease
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
