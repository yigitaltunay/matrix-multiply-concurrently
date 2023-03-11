package utils

import (
	"math/rand"
	"time"
)

func CreateMatrix(rows, cols int) [][]int {
	rand.Seed(time.Now().UnixNano())

	// Create a matrix of the given size
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Fill the matrix with random numbers between 0 and 99
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100)
		}
	}
	return matrix
}
