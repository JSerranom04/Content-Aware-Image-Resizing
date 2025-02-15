package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if len(os.Args) != 4 {
		fmt.Println("Usage: costuras <number_of_seams> <input_file> <results_directory>")
		os.Exit(1)
	}

	seamNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: the number of seams must be an integer")
		os.Exit(1)
	}

	fileIn := os.Args[2]
	resultsDir := os.Args[3]

	// Create the results directory if it doesn't exist
	if err := os.MkdirAll(resultsDir, 0755); err != nil {
		fmt.Printf("Error creating results directory: %v\n", err)
		os.Exit(1)
	}

	// Read the initial image
	log.Printf("Reading image: %s", fileIn)
	imageMatrix := readImage(fileIn)
	log.Printf("Image dimensions: %dx%d", len(imageMatrix), len(imageMatrix[0]))

	// Process each seam
	for i := 0; i < seamNumber; i++ {
		log.Printf("Processing seam %d of %d", i+1, seamNumber)
		
		energyMatrix := EcuRecurrencyMatrix(imageMatrix)
		log.Printf("Energy matrix calculated")
		
		seam := FindMinSeam(energyMatrix)
		log.Printf("Seam found: %v", seam)
		
		imageMatrix = RemoveSeam(imageMatrix, seam)
		log.Printf("Seam removed. New dimensions: %dx%d", 
			len(imageMatrix), len(imageMatrix[0]))
		
		intermediateFile := filepath.Join(resultsDir, fmt.Sprintf("resultado_%d.png", i+1))
		writeImage(intermediateFile, imageMatrix)
		log.Printf("Intermediate image saved: %s", intermediateFile)
	}
}
