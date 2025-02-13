package main

// Returns the lowest number
func Min(a int, b int, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	}
	return c
}

// Recurrence function to calculate seam value
// c(i, j) = min{c(i - 1, j - 1), c(i - 1, j), c(i - 1, j + 1)} + e(i, j)
// c(i, 1) = e(i, 1)
func EcuRecurrency(i int, j int, memory [][]int, Image [][]MatrixComponent) int {
	if memory[i][j] != 0 {
		return memory[i][j]
	}
	ePixel := PixelEnergy(i, 1, Image)
	if j == 1 {
		return ePixel
	}
	// We calcule the minimum out of the range of the cut
	upleft := EcuRecurrency(i-1, j-1, memory, Image)
	left := EcuRecurrency(i-1, j, memory, Image)
	botomleft := EcuRecurrency(i-1, j+1, memory, Image)
	min := min(upleft, left, botomleft)

	// We assign the minimum to the component of the new matrix
	memory[i][j] = min
	return min
}

func EcuRecurrencyMatrix(Image [][]MatrixComponent) [][]MatrixComponent {
	N := len(Image)
	M := len(Image)
	var memory [][]int
	var i int
	for int i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			EcuRecurrency(i, j, memory, Image)
		}
	}
	return Image
}
