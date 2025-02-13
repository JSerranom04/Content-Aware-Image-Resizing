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

// Having this matrix
// A B C
// D E F
// G H I
// energiax = a + 2d + g - c - 2f - i
// energiay = a + 2b + c - g - 2h - i
// energy = sqrt(energiax²+energiay²)
// Returns the energy of the pixel at coordinates (i, j)
func PixelEnergy(i int, j int, Image [][]MatrixComponent) int {
	xenergy := Image[i-1][j-1].brightness + 2*Image[i][j-1].brightness + Image[i+1][j-1].brightness - Image[i-1][j+1].brightness - 2*Image[i][j+1].brightness - Image[i+1][j+1].brightness
	yenergy := Image[i-1][j-1].brightness + 2*Image[i-1][j].brightness + Image[i-1][j+1].brightness - Image[i+1][j-1].brightness - 2*Image[i+1][j].brightness - Image[i+1][j+1].brightness
	energy := int(math.Sqrt(float64(xenergy*xenergy + yenergy*yenergy)))

	return energy
}

// Returns the brightness of the pixel at coordinates (i, j)
func PixelBrightness(i int, j int, image [][]MatrixComponent) int {
	return image[i][j].r + image[i][j].g + image[i][j].b
}
