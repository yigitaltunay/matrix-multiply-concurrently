package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/yigitaltunay/matrix-multiply-concurrently/data"
	matrixmultiply "github.com/yigitaltunay/matrix-multiply-concurrently/matrix-multiply"
	"github.com/yigitaltunay/matrix-multiply-concurrently/utils"
)

func main() {

	rows, cols := 100, 100

	data.Matrix1 = utils.CreateMatrix(rows, cols)
	data.Matrix2 = utils.CreateMatrix(rows, cols)

	t1 := time.Now()
	result1 := matrixmultiply.GeneralMulply()
	color.Blue("General Mulpy : ")
	fmt.Println(time.Until(t1))

	t1 = time.Now()
	result2 := matrixmultiply.ConcurrentlyMultiply()
	color.Blue("Concurrently Mulpy : ")
	fmt.Println(time.Until(t1))

	color.Green("Matrices Equal : %dx%d", rows, cols)

	fmt.Println(utils.MatricesEqual(result1, result2))
}
