package faii

func flipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		n := len(v)
		if n%2 == 1 {
			v[n/2] ^= 1
		}
		for i := 0; i < n/2; i++ {
			tmp := v[i]^1
			v[i] = v[n-i-1]^1
			v[n-i-1] = tmp
		}
	}
	return A
}
