package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	start := 3
	end := 5
	count := 5

	runTimes := map[int][]int{}

	for i := start; i <= end; i++ {
		runTimes[i] = make([]int, count)
		cmdStr := fmt.Sprintf("cd ../1st/; go run main.go %d", i)

		fmt.Printf("%d: ", i)

		for j := 0; j < count; j++ {
			out, err := exec.Command("sh", "-c", cmdStr).Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			time, err := strconv.Atoi(strings.ReplaceAll(string(out), "\n", ""))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			runTimes[i][j] = time
			fmt.Printf("%d ", time)
		}

		fmt.Printf("\n")
	}

}
