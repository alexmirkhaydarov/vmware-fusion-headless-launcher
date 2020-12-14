package main

import (
	"os"
	"fmt"
	"os/exec"
	"log"
	"io/ioutil"
	"strconv"
)

func main() {
	// Checking if VMware is installed
	if _, err := os.Stat("/Applications/VMware Fusion.app"); os.IsNotExist(err) {
		log.Fatal(err)
	}

	// Checking if vmrun is available
	cmd := exec.Command("vmrun")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(homeDir + "/Virtual Machines.localized")

	if _, err := os.Stat(homeDir + "/Virtual Machines.localized"); os.IsNotExist(err) {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(homeDir + "/Virtual Machines.localized")
	if err != nil {
		log.Fatal(err)
	}

	for i, f := range files {
		if f.Name() != ".DS_Store" && f.Name() != ".localized" {
			fmt.Println(strconv.Itoa(i) + " " + f.Name())
		}
	}
}
