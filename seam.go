package main

// FindMinSeam finds the minimum seam in the energy matrix
func FindMinSeam(energyMatrix [][]int) []int {
	N := len(energyMatrix)
	M := len(energyMatrix[0])
	seam := make([]int, N)

	// Find the minimum in the last row
	minVal := energyMatrix[N-1][0]
	minIdx := 0
	for j := 1; j < M; j++ {
		if energyMatrix[N-1][j] < minVal {
			minVal = energyMatrix[N-1][j]
			minIdx = j
		}
	}
	seam[N-1] = minIdx

	// Reconstruct the path upwards
	for i := N - 2; i >= 0; i-- {
		j := seam[i+1]
		minVal = energyMatrix[i][j]
		minIdx = j

		// Check the three possible upper neighbors
		if j > 0 && energyMatrix[i][j-1] < minVal {
			minVal = energyMatrix[i][j-1]
			minIdx = j - 1
		}
		if j < M-1 && energyMatrix[i][j+1] < minVal {
			minVal = energyMatrix[i][j+1]
			minIdx = j + 1
		}
		seam[i] = minIdx
	}

	return seam
}

// RemoveSeamFromImage removes a vertical seam from the image
func RemoveSeamFromImage(image [][]MatrixComponent, seam []int) [][]MatrixComponent {
	N := len(image)
	M := len(image[0])
	newImage := make([][]MatrixComponent, N)

	for i := 0; i < N; i++ {
		newImage[i] = make([]MatrixComponent, M-1)
		seamJ := seam[i]
		idx := 0
		for j := 0; j < M; j++ {
			if j != seamJ {
				newImage[i][idx] = image[i][j]
				idx++
			}
		}
	}

	return newImage
}
