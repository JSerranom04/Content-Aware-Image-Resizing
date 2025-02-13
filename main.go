package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func readImage(name string) [][]MatrixComponent {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening image", err)
		os.Exit(1)
	}

	data, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image", err)
		os.Exit(1)
	}

	N := data.Bounds().Max.X - data.Bounds().Min.X
	M := data.Bounds().Max.Y - data.Bounds().Min.Y
	imageRet := make([][]MatrixComponent, N)
	for i := range imageRet {
		imageRet[i] = make([]MatrixComponent, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			r, g, b, _ := data.At(i, j).RGBA()
			imageRet[i][j].r = int(r)
			imageRet[i][j].g = int(g)
			imageRet[i][j].b = int(b)
			imageRet[i][j].brightness = int(r + g + b)
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

func main() {

	if len(os.Args) != 4 {
		fmt.Println("LO HAS EJECUTADO MAL, ESTUPIDO!")
	}
	//seamNumber := os.Args[1]
	fileIn := os.Args[2]
	//fileOut := os.Args[3]

	// We open the image in order to read the matrix and convert it to the image format
	// [i,j] (r,g,b,r+g+b)
	Image := readImage(fileIn)
	//readImage(fileIn)

	fmt.Println(Image)
}
