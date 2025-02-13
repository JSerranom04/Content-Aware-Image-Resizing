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
	var imageRet [][]MatrixComponent
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
		}
	}
	return imageRet
}

func printImage(img image.Image) {
	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
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
	seamNumber := os.Args[1]
	fileIn := os.Args[2]
	fileOut := os.Args[3]

	// We open the image in order to read the matrix and convert it to the image format
	// [i,j] (r,g,b,r+g+b)
	Image := readImage(fileIn)

	var recurrencyTable [][]int

	fmt.Println(Image, recurrencyTable, seamNumber, fileIn, fileOut)
}
