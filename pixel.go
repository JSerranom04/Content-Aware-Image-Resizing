package main

import (
	"math"
	"runtime"
	"sync"
)

// Assigns the new brightness to the components of the matrix
func CalculateBrightnesOfImage(image [][]MatrixComponent) [][]MatrixComponent {
	N := len(image)
	M := len(image[0])
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	
	// Dividir el trabajo en varias goroutines
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			// Cada goroutine procesa una parte de las filas
			for i := workerId; i < N; i += numWorkers {
				for j := 0; j < M; j++ {
					image[i][j].brightness = PixelBrightness(i, j, image)
				}
			}
		}(w)
	}
	wg.Wait()
	return image
}

// Gives the brightness of a given pixel and prevents it from returning
// a null value when the component does not exist
func getValue(x, y int, Image [][]MatrixComponent) int {
	if x < 0 || y < 0 || x >= len(Image) || y >= len(Image[0]) {
		return 0 // Si está fuera de los límites, devuelve 0
	}
	return Image[x][y].brightness
}

// Gives the value of the given matrix components in variables after checking if
// they are null and if so, it returns a 0 in that value
func GiveMeTheMatrixComponents(i int, j int, Image [][]MatrixComponent) (int, int, int, int, int, int, int, int) {
	a := getValue(i-1, j-1, Image)
	b := getValue(i-1, j, Image)
	c := getValue(i-1, j+1, Image)
	d := getValue(i, j-1, Image)
	f := getValue(i, j+1, Image)
	g := getValue(i+1, j-1, Image)
	h := getValue(i+1, j, Image)
	ii := getValue(i+1, j+1, Image) // Cambié el nombre para evitar conflicto con `i` del parámetro

	return a, b, c, d, f, g, h, ii
}

// Having this matrix
// A B C
// D E F
// G H I
// energiax = a + 2d + g - c - 2f - i
// energiay = a + 2b + c - g - 2h - i
// energy = sqrt(energiax²+energiay²)
// Returns the energy of the pixel at coordinates (i, j)
func PixelEnergy(i int, j int, Image [][]MatrixComponent) int {
	a, b, c, d, f, g, h, k := GiveMeTheMatrixComponents(i, j, Image)
	xenergy := a + 2*d + g - c - 2*f - k
	yenergy := a + 2*b + c - g - 2*h - k
	energy := int(math.Sqrt(float64(xenergy*xenergy + yenergy*yenergy)))

	return energy
}

// Returns the brightness of the pixel at coordinates (i, j)
// Modular in case we want to change the brightness formula
func PixelBrightness(i int, j int, image [][]MatrixComponent) int {
	return image[i][j].r + image[i][j].g + image[i][j].b
}

// Calculates the energy value for all pixels in "image" and stores it in "image"
func calculateEnergyOfImageInitial(image [][]MatrixComponent) {
	N := len(image)
	M := len(image[0])
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU() // Usar número de CPU disponibles
	
	// Dividir el trabajo en varias goroutines
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			// Cada goroutine procesa una parte de las filas
			for i := workerId; i < N; i += numWorkers {
				for j := 0; j < M; j++ {
					image[i][j].energy = PixelEnergy(i, j, image)
				}
			}
		}(w)
	}
	wg.Wait()
}

// Calculates the energy value for all pixels that are touching the "seam" and stores it in "image"
func calculateEnergyOfImage(image [][]MatrixComponent, seam []int) {
	M := len(image[0])
	var wg sync.WaitGroup
	
	for i, v := range seam {
		wg.Add(1)
		go func(i, v int) {
			defer wg.Done()
			// Left pixel of the one removed
			if v > 0 {
				image[i][v-1].energy = PixelEnergy(i, v-1, image)
			}
			// Right pixel of the one removed
			if v < M-1 {
				image[i][v].energy = PixelEnergy(i, v, image)
			}
		}(i, v)
	}
	wg.Wait()
}
