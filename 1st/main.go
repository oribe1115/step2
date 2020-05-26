package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

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

	start := time.Now()
	c := lib.MultiplyMatrices(a, b)
	end := time.Now()

	total := lib.SumMatrixElements(c)

	if len(args) > 1 && args[1] == "log" {
		fmt.Println("A")
		lib.PrintMatrix(a)
		fmt.Println("B")
		lib.PrintMatrix(b)

		fmt.Println("---")

		fmt.Println("C")
		lib.PrintMatrix(c)
		fmt.Printf("total: %d\n", total)

		fmt.Println("---")

		fmt.Println("実行時間(ナノ秒)")
	}

	fmt.Println(end.Sub(start).Nanoseconds())

}
