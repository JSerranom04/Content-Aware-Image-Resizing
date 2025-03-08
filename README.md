# Dinamic programming - Seam Carving
Asignatura: Algoritmia Básica
Curso 2023/24

## Directory structure

### Main files
- `main.go`: Main program that implements the seam carving algorithm
- `pixel.go`: Functions for the calculation of energy and brightness of pixels
- `recurrency.go`: Implementation of dynamic programming
- `seam.go`: Functions to find and remove seams
- `image.go`: Functions to read/write images
- `types.go`: Definition of data types
- `seam_test.go`: Unit tests
- `ejecutar.sh`: Script to automate the execution
- `LEEME.md`: This file

### Requirements
- Go 1.20 or superior
- Standard Go library (no requires external dependencies)

### Compilation

bash
go build -o costuras .go

### Execution

bash
./costuras <number_of_seams> <input_image_file> <results_directory>

Example:

bash
./costuras 50 ./testImages/imagen.png ./results


### Description of the algorithm
The program implements the seam carving algorithm to reduce the width of an image by removing vertical seams of minimum energy. The process consists of three main steps:

1. Energy calculation: For each pixel, its energy is calculated using the gradient in X and Y
2. Dynamic programming: The minimum energy seam is found using the recurrence equation
3. Removal: The seam is removed and the process is repeated

### Data structure
- `MatrixComponent`: Structure that stores the RGBA values and brightness of each pixel
- Two-dimensional matrices to represent the image and the energies

### Tests
To run the tests:

bash
go test -v




### Test images
The `pruebas/` directory contains example images:
- boat.png: Image of a boat
- water.png: Image of water with homogeneous areas
- rocket.png: Image of a rocket
- imagen.png: General test image

### Additional notes
- The program accepts PNG images
- The results are saved in the specified directory
- The progress is shown during the execution
- The execution time depends on the image size and the number of seams
- The implementation includes optimizations to recalculate only the affected areas by each seam

### Performance example
In an image of 1466x1220 pixels, reducing 350 seams:
- Initial dimensions: 1466x1220
- Final dimensions: 1466x870
- Approximate time: 2-3 minutes

Authors: Victor Orrios Baron and Juan José Serrano Mora