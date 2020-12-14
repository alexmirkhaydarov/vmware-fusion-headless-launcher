package main

import (
	"path/filepath"
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

	if _, err := os.Stat(homeDir + "/Virtual Machines.localized"); os.IsNotExist(err) {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(homeDir + "/Virtual Machines.localized")
	if err != nil {
		log.Fatal(err)
	}

	for i, f := range files {
		if filepath.Ext(f.Name()) == ".vmwarevm" {
			fmt.Println(strconv.Itoa(i - 1) + " " + f.Name())
		}
	}
}
