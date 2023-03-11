package matrixmultiply

import (
	"fmt"

	"github.com/yigitaltunay/matrix-multiply-concurrently/data"
)

func GeneralMulply() [][]int {
	// Define two matrices
	matrix1 := data.Matrix1
	matrix2 := data.Matrix2

	// Get dimensions of matrices
	m1Rows := len(matrix1)
	m1Cols := len(matrix1[0])
	m2Rows := len(matrix2)
	m2Cols := len(matrix2[0])

	// Check if matrices can be multiplied
	if m1Cols != m2Rows {
		fmt.Println("Matrices cannot be multiplied")
		return [][]int{}
	}

	// Initialize the resulting matrix with zeros
	result := make([][]int, m1Rows)
	for i := range result {
		result[i] = make([]int, m2Cols)
	}

	// Multiply matrices
	for i := 0; i < m1Rows; i++ {
		for j := 0; j < m2Cols; j++ {
			for k := 0; k < m1Cols; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
	// Print the result
	//	fmt.Println(result)
}
