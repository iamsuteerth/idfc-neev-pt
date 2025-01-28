package questions

func Triangle(n int) [][]int {
	var pascal [][]int
	row1 := []int{1}
	row2 := []int{1, 1}
	if n == 1 {
		return append(pascal, row1)
	}
	if n == 2 {
		return append(pascal, row1, row2)
	}
	pascal = append(pascal, row1, row2)
	for i := 3; i <= n; i++ {
		curr := make([]int, i)
		for j := 0; j < len(pascal[i-2])-1; j++ {
			curr[j+1] = pascal[i-2][j] + pascal[i-2][j+1]
		}
		curr[0] = 1
		curr[len(curr)-1] = 1
		pascal = append(pascal, curr)
	}
	return pascal
}
