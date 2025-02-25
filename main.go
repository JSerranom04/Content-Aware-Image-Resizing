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
	N := len(imageMatrix)
	M := len(imageMatrix[0])

	// Initial calculation of energy
	calculateEnergyOfImageInitial(imageMatrix)
	log.Printf("Energy matrix calculated")

	// Initial calculation of recurrency matrix
	recurrencyMatrix := EcuRecurrencyMatrixInitial(imageMatrix)
	log.Printf("Recurrency matrix calculated")

	// Process all seams
	for i := 0; i < seamNumber; i++ {
		log.Printf("Processing seam %d of %d", i+1, seamNumber)

		seam := FindMinSeam(recurrencyMatrix)

		imageMatrix = RemoveSeamFromImage(imageMatrix, seam)
		//log.Printf("Seam removed. Current dimensions: %dx%d", len(imageMatrix), len(imageMatrix[0]))

		if i < seamNumber-1 {
			// Recalculate the energy of the pixels that had their neighbours changed
			calculateEnergyOfImage(imageMatrix, seam)
			// Recalculate the recurrency value for all pixels under the piramid
			recurrencyMatrix = EcuRecurrencyMatrix(recurrencyMatrix, seam, imageMatrix)
		}
	}

	// Save final image after all seams have been removed
	finalFile := filepath.Join(resultsDir, "resultado_final.png")
	writeImage(finalFile, imageMatrix)
	log.Printf("Final image saved: %s", finalFile)
	log.Printf("Initial image dimensions: %dx%d", N, M)
	log.Printf("Final dimensions: %dx%d", len(imageMatrix), len(imageMatrix[0]))
}
