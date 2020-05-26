package main

import (
	"fmt"

	"github.com/oribe1115/step2/1st-graph/lib"
)

func main() {
	start := 3
	end := 5
	count := 5

	_, err := lib.MeasureRunTimes(start, end, count)
	if err != nil {
		fmt.Println(err)
	}

}
