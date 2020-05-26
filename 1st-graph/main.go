package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmdStr := "cd ../1st/; go run main.go 3"
	out, err := exec.Command("sh", "-c", cmdStr).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
}
