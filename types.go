package main

// This type defines the componentes assigned to each component of the matrix
// It stores NRGBA values and brightness of pixel (r+g+b)
type MatrixComponent struct {
	r, g, b, a int
	brightness int
}
