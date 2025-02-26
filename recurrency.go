package main

import (
	"runtime"
	"sync"
)

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
	ePixel := Image[i][j].energy
	// Base case, pixel is on top row
	if i == 0 {
		return ePixel
	}
	// We calculate the minimum out of the range of the cut
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
	// Setting bounds
	N := len(Image)
	M := len(Image[0])
	// Initialize the memory matrix
	memory := make([][]int, N)
	for i := 0; i < N; i++ {
		memory[i] = make([]int, M)
		for j := 0; j < M; j++ {
			memory[i][j] = -1
		}
	}
	
	// Calculate first row (base case)
	for j := 0; j < M; j++ {
		memory[0][j] = Image[0][j].energy
	}
	
	// For the rest of the rows, we use row-wise concurrency
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	
	// Start from row 1 since row 0 is the base case
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			// Each goroutine processes a set of rows
			for i := 1 + workerId; i < N; i += numWorkers {
				for j := 0; j < M; j++ {
					EcuRecurrency(i, j, &memory, Image)
				}
			}
		}(w)
	}
	wg.Wait()
	
	return memory
}

func EcuRecurrencyMatrix(matrix [][]int, seam []int, Image [][]MatrixComponent) [][]int {
	// Setting bounds
	N := len(Image)
	M := len(Image[0])
	// Initialize the memory matrix and copy non affected pixel from the original recurrency matrix
	memory := make([][]int, N)
	for i := 0; i < N; i++ {
		memory[i] = make([]int, M)
		for j := 0; j < M; j++ {
			memory[i][j] = -1
		}
		// Affected pixels are the ones inside the pyramid generated from removing the top pixel
		// If image is square this saves calculating half of the recurrency matrix at least
		copy(memory[i][:max(0, seam[0]-i)], matrix[i][:max(0, seam[0]-i)])
		copy(memory[i][min(M, seam[0]+1+i):M], matrix[i][min(M, seam[0]+1+i):M])
	}
	
	// Calculate first row (base case) for the affected region
	for j := max(0, seam[0]-0); j < min(M, seam[0]+1+0); j++ {
		memory[0][j] = Image[0][j].energy
	}
	
	// Calculate the rest of the rows in parallel
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			// Each goroutine processes a set of rows
			for i := 1 + workerId; i < N; i += numWorkers {
				// Only process the affected region by the seam
				for j := max(0, seam[0]-i); j < min(M, seam[0]+1+i); j++ {
					if memory[i][j] == -1 { // If it hasn't been copied from the original
						EcuRecurrency(i, j, &memory, Image)
					}
				}
			}
		}(w)
	}
	wg.Wait()
	
	return memory
}
