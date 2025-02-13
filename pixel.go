package main

import "math"

// Assigns the new bright to the components of the matrix
func CalculateBrightnesOfImage(image [][]MatrixComponent) [][]MatrixComponent {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image); j++ {
			// Call the function that calculates the brightness
			image[i][j].brightness = PixelBrightness(i, j, image)
		}
	}

	return image
}
func getValue(x, y int, Image [][]MatrixComponent) int {
	if x < 0 || y < 0 || x >= len(Image) || y >= len(Image[0]) {
		return 0 // Si está fuera de los límites, devuelve 0
	}
	return Image[x][y].brightness
}

func GiveMeTheMatrixComponents(i int, j int, Image [][]MatrixComponent) (int, int, int, int, int, int, int, int, int) {

	a := getValue(i-1, j-1)
	b := getValue(i-1, j)
	c := getValue(i-1, j+1)
	d := getValue(i, j-1)
	e := getValue(i, j)
	f := getValue(i, j+1)
	g := getValue(i+1, j-1)
	h := getValue(i+1, j)
	ii := getValue(i+1, j+1) // Cambié el nombre para evitar conflicto con `i` del parámetro

	return a, b, c, d, e, f, g, h, ii
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
	a, b, c, d, e, f, g, h, i := GiveMeTheMatrixComponents(i, j, Image)
	xenergy := a + 2*d + g - c - 2*f - i
	yenergy := a + 2*b + c - g - 2*h - i
	energy := int(math.Sqrt(float64(xenergy*xenergy + yenergy*yenergy)))

	return energy
}

// Returns the brightness of the pixel at coordinates (i, j)
func PixelBrightness(i int, j int, image [][]MatrixComponent) int {
	return image[i][j].r + image[i][j].g + image[i][j].b
}
