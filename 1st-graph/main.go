package main

import (
	"fmt"

	"github.com/oribe1115/step2/1st-graph/lib"
)

func main() {
	start := 3
	end := 10
	count := 10

	runTimes, err := lib.MeasureRunTimes(start, end, count)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := lib.SetData(runTimes, start, end)

	err = lib.PlotGraph(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
