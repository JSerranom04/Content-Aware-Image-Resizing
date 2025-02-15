package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Reads a PNG from the file "name" and return the NRGBA+Bightnes values of the image
func readImage(name string) [][]MatrixComponent {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Error opening image", err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image", err)
		os.Exit(1)
	}

	N := data.Bounds().Max.X - data.Bounds().Min.X
	M := data.Bounds().Max.Y - data.Bounds().Min.Y
	imageRet := make([][]MatrixComponent, M)

	for i := 0; i < M; i++ {
		imageRet[i] = make([]MatrixComponent, N)
		for j := 0; j < N; j++ {
			var r, g, b, a uint8

			// Intentar convertir a NRGBA primero
			if nrgba, ok := data.At(j, i).(color.NRGBA); ok {
				r, g, b, a = nrgba.R, nrgba.G, nrgba.B, nrgba.A
			} else {
				// Si no es NRGBA, usar el método RGBA() y convertir
				r32, g32, b32, a32 := data.At(j, i).RGBA()
				r = uint8(r32 >> 8)
				g = uint8(g32 >> 8)
				b = uint8(b32 >> 8)
				a = uint8(a32 >> 8)
			}

			imageRet[i][j].r = int(r)
			imageRet[i][j].g = int(g)
			imageRet[i][j].b = int(b)
			imageRet[i][j].a = int(a)
			imageRet[i][j].brightness = int(r) + int(g) + int(b)
		}
	}
	return imageRet
}

// Taken from https://golangdocs.com/golang-image-processing
func printImage(data image.Image) {
	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := data.Bounds().Min.Y; y < data.Bounds().Max.Y; y++ {
		for x := data.Bounds().Min.X; x < data.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(data.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}

// Given the new file to write "name" and the NRGBA values of the image it creates/overrides the new image
func writeImage(name string, imageMatrix [][]MatrixComponent) {
	// imageMatrix must be a rectangular 2d matrix
	N := len(imageMatrix)
	M := len(imageMatrix[0])
	newBounds := image.Rect(0, 0, M, N)
	img := image.NewNRGBA(newBounds)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			// It writes the unprocesed NRGBA values with a 100% lossless write
			img.Set(j, i, color.NRGBA{
				R: uint8(imageMatrix[i][j].r),
				G: uint8(imageMatrix[i][j].g),
				B: uint8(imageMatrix[i][j].b),
				A: uint8(imageMatrix[i][j].a),
			})
		}
	}
	newImageFile, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}

	defer newImageFile.Close()
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}
	err = png.Encode(newImageFile, img)
	if err != nil {
		newImageFile.Close()
		fmt.Println("Error encoding new file", err)
		os.Exit(1)
	}

}
