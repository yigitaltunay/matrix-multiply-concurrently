# Golang Matrix Multiply Concurrently





```go
func main() {

	rows, cols := 50000, 50000

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

```

```go
General Mulpy : 
-15m12.05931s
Concurrently Mulpy : 
-42.04266s
Matrices Equal : 5000x5000
true
```


