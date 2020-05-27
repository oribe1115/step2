package main

import (
	"fmt"

	"github.com/oribe1115/step2/4th/lib"
)

func main() {

	list := lib.InitHistoryList(4)
	list.Print()
	list.Cache("a")
	list.Print()
	list.Cache("b")
	list.Print()
	list.Cache("c")
	list.Print()
	list.Cache("d")
	list.Print()

	fmt.Println()

	list.Cache("d")
	list.Print()
	list.Cache("a")
	list.Print()
	list.Cache("c")
	list.Print()
	list.Cache("a")
	list.Print()
	list.Cache("a")
	list.Print()

	fmt.Println()

	list.Cache("e")
	list.Print()
	list.Cache("f")
	list.Print()
}
