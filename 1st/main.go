package main

import (
	"fmt"

	"github.com/oribe1115/step2/1st/lib"
)

func main() {
	a, b := lib.InitMatrices(5)
	fmt.Println("A")
	lib.PrintMatrix(a)
	fmt.Println("B")
	lib.PrintMatrix(b)

	c := lib.MultiplyMatrices(a, b)
	fmt.Println("---")
	fmt.Println("C")
	lib.PrintMatrix(c)

	total := lib.SumMatrixElements(c)
	fmt.Println(total)
}
