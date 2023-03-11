package utils

func MatricesEqual(a, b [][]int) bool {
	// Check if the matrices have the same dimensions
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}

	// Check if the elements in the matrices are equal
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	// The matrices are equal
	return true
}
