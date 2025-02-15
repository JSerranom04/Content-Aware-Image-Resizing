package main

import (
	"reflect"
	"testing"
)

func TestFindMinSeam(t *testing.T) {
    tests := []struct {
        name         string
        energyMatrix [][]int
        want         []int
    }{
        {
            name: "Matriz 3x3 básica",
            energyMatrix: [][]int{
                {1, 4, 3},
                {8, 2, 6},
                {3, 9, 2},
            },
            want: []int{0, 1, 2},
        },
        {
            name: "Matriz con camino claro",
            energyMatrix: [][]int{
                {9, 1, 9},
                {9, 2, 9},
                {9, 1, 9},
            },
            want: []int{1, 1, 1},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FindMinSeam(tt.energyMatrix)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("FindMinSeam() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRemoveSeam(t *testing.T) {
    image := [][]MatrixComponent{
        {{r: 1, g: 1, b: 1}, {r: 2, g: 2, b: 2}, {r: 3, g: 3, b: 3}},
        {{r: 4, g: 4, b: 4}, {r: 5, g: 5, b: 5}, {r: 6, g: 6, b: 6}},
    }
    seam := []int{1, 1}

    result := RemoveSeam(image, seam)

    // Verificar dimensiones
    if len(result) != 2 || len(result[0]) != 2 {
        t.Errorf("Dimensiones incorrectas después de RemoveSeam: %dx%d", len(result), len(result[0]))
    }

    // Verificar valores
    expected := [][]MatrixComponent{
        {{r: 1, g: 1, b: 1}, {r: 3, g: 3, b: 3}},
        {{r: 4, g: 4, b: 4}, {r: 6, g: 6, b: 6}},
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("RemoveSeam() = %v, want %v", result, expected)
    }
}

func TestEcuRecurrencyMatrix(t *testing.T) {
    image := [][]MatrixComponent{
        {{brightness: 1}, {brightness: 4}, {brightness: 3}},
        {{brightness: 8}, {brightness: 2}, {brightness: 6}},
        {{brightness: 3}, {brightness: 9}, {brightness: 2}},
    }

    result := EcuRecurrencyMatrix(image)

    if len(result) != 3 || len(result[0]) != 3 {
        t.Errorf("Dimensiones incorrectas de la matriz de energía: %dx%d", len(result), len(result[0]))
    }
} 