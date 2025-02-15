package main

import "fmt"

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
	if (*memory)[i][j] != 0 {
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

func EcuRecurrencyMatrix(Image [][]MatrixComponent) [][]int {
	// Seting bounds
	N := len(Image)
	M := len(Image[0])
	// Initialize the memomory matrix
	memory := make([][]int, N)
	for i := 0; i < N; i++ {
		memory[i] = make([]int, M)
	}
	// Calculate recurrency values for the whole image
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Println("We are in the recurrence function, i = ", i, " j = ", j)
			EcuRecurrency(i, j, &memory, Image)
			fmt.Println("We are out of the recurrence function, i = ", i, " j = ", j)
		}
	}
	return memory
}
