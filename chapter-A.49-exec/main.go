package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output1, _ := exec.Command("ls", "-lah").Output()
	output2, _ := exec.Command("pwd").Output()
	output3, _ := exec.Command("whoami").Output()

	fmt.Println(string(output1))
	fmt.Println(string(output2))
	fmt.Println(string(output3))
}
