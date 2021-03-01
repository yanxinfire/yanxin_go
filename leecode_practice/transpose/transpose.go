package transpose

func transpose(matrix [][]int) [][]int {
	n1 := len(matrix)
	n2 := len(matrix[0])
	res := [][]int{}
	for i := 0; i < n2; i++ {
		r := make([]int, 0, n1)
		for q := 0; q < n1; q++ {
			r = append(r, matrix[q][i])
		}
		res = append(res, r)
	}
	return res
}
