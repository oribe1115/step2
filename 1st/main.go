package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/oribe1115/step2/1st/lib"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("行列のサイズを指定する引数を入力してください")
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	a, b := lib.InitMatrices(n)
	fmt.Println("A")
	lib.PrintMatrix(a)
	fmt.Println("B")
	lib.PrintMatrix(b)

	c := lib.MultiplyMatrices(a, b)
	fmt.Println("---")
	fmt.Println("C")
	lib.PrintMatrix(c)

	total := lib.SumMatrixElements(c)
	fmt.Println("---")
	fmt.Println(total)
}
