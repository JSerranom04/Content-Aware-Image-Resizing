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
func EcuRecurrency(i int, j int, memory *[][]int, Image [][]MatrixComponent) int {
	// If the pixel recurrency function has already been calculated
	if (*memory)[i][j] != -1 {
		return (*memory)[i][j]
	}
	ePixel := PixelEnergy(i, j, Image)
	// Base case, pixel is on top row
	if i == 0 {
		return ePixel
	}
	// We calcule the minimum out of the range of the cut
	M := len(Image[0])
	var upleft, left, bottomleft int
	if j > 0 {
		upleft = EcuRecurrency(i-1, j-1, memory, Image)
	} else {
		upleft = int(^uint(0) >> 1) // Max int value
	}
	left = EcuRecurrency(i-1, j, memory, Image)
	if j < M-1 {
		bottomleft = EcuRecurrency(i-1, j+1, memory, Image)
	} else {
		bottomleft = int(^uint(0) >> 1) // Max int value
	}
	min := Min(upleft, left, bottomleft)

	// We assign the minimum to the component of the new matrix
	(*memory)[i][j] = min + ePixel
	return min + ePixel
}

func EcuRecurrencyMatrixInitial(Image [][]MatrixComponent) [][]int {
	// Seting bounds
	N := len(Image)
	M := len(Image[0])
	// Initialize the memomory matrix
	memory := make([][]int, N)
	for i := 0; i < N; i++ {
		memory[i] = make([]int, M)
		for j := 0; j < M; j++ {
			memory[i][j] = -1
		}
	}
	// Calculate recurrency values for the whole image
	for i := 0; i < N; i++ {
		//fmt.Println("Calculating recurrency matrix:", i, "/", N)
		for j := 0; j < M; j++ {
			EcuRecurrency(i, j, &memory, Image)
		}
	}
	return memory
}

func EcuRecurrencyMatrix(matrix [][]int, seam []int, Image [][]MatrixComponent) [][]int {
	// Seting bounds
	N := len(Image)
	M := len(Image[0])
	// Initialize the memomory matrix
	memory := make([][]int, N)
	for i := 0; i < N; i++ {
		memory[i] = make([]int, M)
		for j := 0; j < M; j++ {
			memory[i][j] = -1
		}
		copy(memory[i][:max(0, seam[0]-i)], matrix[i][:max(0, seam[0]-i)])
		copy(memory[i][min(M, seam[0]+1+i):M], matrix[i][min(M, seam[0]+1+i):M])
	}
	// Calculate recurrency values for the whole image
	for i := 0; i < N; i++ {
		//fmt.Println("Calculating recurrency matrix:", i, "/", N)
		for j := max(0, seam[0]-i); j < min(M, seam[0]+1+i); j++ {
			EcuRecurrency(i, j, &memory, Image)
		}
	}
	return memory
}
