package questions

func SpiralMatrix(size int) [][]int {
	xDir := []int{0, 1, 0, -1}
	yDir := []int{1, 0, -1, 0}

	r := 0
	c := 0
	step := 0

	res := make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
	}

	for i := 0; i < size*size; i++ {
		res[r][c] = i + 1
		newR, newC := r+xDir[step], c+yDir[step]
		if newR < 0 || newR >= size || newC < 0 || newC >= size || res[newR][newC] != 0 {
			step = (step + 1) % 4 // Change direction
		}
		r += xDir[step]
		c += yDir[step]
	}
	return res
}
