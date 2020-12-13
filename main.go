package main

import (
	"os"
	"fmt"
	"os/exec"
	// "log"
)

func main() {
	// Checking if VMware is installed
	if _, err := os.Stat("/Applications/VMware Fusion.app"); os.IsNotExist(err) {
		fmt.Println(err)
	}

	// Checking if vmrun is in the PATH
	cmd := exec.Command("vmrun")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}
