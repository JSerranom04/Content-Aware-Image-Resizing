package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("seamNumber fileInName fileOutName")
		os.Exit(1)
	}
	seamNumber, err := strconv.Atoi(os.Args[1])
	println(seamNumber)
	if err != nil {
		fmt.Println("seamNumber fileInName fileOutName")
		os.Exit(1)
	}
	fileIn := os.Args[2]
	fileOut := os.Args[3]
	fmt.Println("Number of seams:", seamNumber)
	fmt.Println("File IN:", fileIn)
	fmt.Println("File OUT:", fileOut)
	imageMatrix := readImage(fileIn)
	writeImage(fileOut, imageMatrix)

}
