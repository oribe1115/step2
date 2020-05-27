package main

import (
	"fmt"

	"github.com/oribe1115/step2/4th/lib"
)

func main() {

	cacheTable := lib.InitCacheTable(4)

	cacheTable.Print()
	cacheTable.Cache("a", "AAA")
	cacheTable.Print()
	cacheTable.Cache("b", "BBB")
	cacheTable.Print()
	cacheTable.Cache("c", "CCC")
	cacheTable.Print()
	cacheTable.Cache("d", "DDD")
	cacheTable.Print()

	fmt.Println()

	cacheTable.Cache("d", "DDD")
	cacheTable.Print()
	cacheTable.Cache("a", "AAA")
	cacheTable.Print()
	cacheTable.Cache("c", "CCC")
	cacheTable.Print()
	cacheTable.Cache("a", "AAA")
	cacheTable.Print()
	cacheTable.Cache("a", "AAA")
	cacheTable.Print()

	fmt.Println()

	cacheTable.Cache("e", "EEE")
	cacheTable.Print()
	cacheTable.Cache("f", "FFF")
	cacheTable.Print()
}
