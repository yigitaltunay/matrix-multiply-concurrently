package matrixmultiply

import (
	"fmt"
	"sync"

	"github.com/yigitaltunay/matrix-multiply-concurrently/data"
)

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {
	m1Rows := len(matrix1)
	m1Cols := len(matrix1[0])
	m2Rows := len(matrix2)
	m2Cols := len(matrix2[0])

	if m1Cols != m2Rows {
		fmt.Println("Matrices cannot be multiplied")
		return nil
	}

	result := make([][]int, m1Rows)
	for i := range result {
		result[i] = make([]int, m2Cols)
	}

	var wg sync.WaitGroup
	wg.Add(m1Rows)

	for i := 0; i < m1Rows; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < m2Cols; j++ {
				for k := 0; k < m1Cols; k++ {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
			}
		}(i)
	}

	wg.Wait()

	return result
}

func ConcurrentlyMultiply() [][]int {
	return multiplyMatrices(data.Matrix1, data.Matrix2)
}
