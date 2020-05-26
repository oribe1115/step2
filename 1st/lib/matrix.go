package lib

import "fmt"

func InitMatrices(n int) ([][]int, [][]int) {
	a := make([][]int, n)
	b := make([][]int, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		b[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = i*n + j
			b[i][j] = j*n + i
		}
	}

	return a, b
}

func PrintMatrix(matrix [][]int) {
	for _, r := range matrix {
		for _, i := range r {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n")
	}
}

func MultiplyMatrices(a [][]int, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = 0
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result
}

func SumMatrixElements(matrix [][]int) int {
	n := len(matrix)
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += matrix[i][j]
		}
	}

	return result
}
