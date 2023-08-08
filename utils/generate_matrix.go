package utils

import (
	"encoding/json"
	"math/rand"
	"os"
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

func SaveMatrixToFile(matrix [][]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(matrix); err != nil {
		return err
	}
	return nil
}

func LoadMatrixFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&matrix); err != nil {
		return nil, err
	}
	return matrix, nil
}
