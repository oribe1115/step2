package lib

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func MeasureRunTimes(start, end, count int) (map[int][]int, error) {
	runTimes := map[int][]int{}

	for i := start; i <= end; i++ {
		runTimes[i] = make([]int, count)
		cmdStr := fmt.Sprintf("cd ../1st/; go run main.go %d", i)

		fmt.Printf("%d: ", i)

		for j := 0; j < count; j++ {
			out, err := exec.Command("sh", "-c", cmdStr).Output()
			if err != nil {
				return nil, err
			}
			time, err := strconv.Atoi(strings.ReplaceAll(string(out), "\n", ""))
			if err != nil {
				return nil, err
			}
			runTimes[i][j] = time
			fmt.Printf("%d ", time)
		}

		fmt.Printf("\n")
	}

	return runTimes, nil
}
