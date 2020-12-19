package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	dir, err := ioutil.ReadDir(homeDir + "/Virtual Machines.localized")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> Select a value")

	for i, f := range dir {
		if filepath.Ext(f.Name()) == ".vmwarevm" {
			fmt.Println(strconv.Itoa(i-1) + " " + f.Name())
		}
	}

	fmt.Println("------------------------------------------")
	var input int
	fmt.Scan(&input)
	// fmt.Println(input)

	for i, f := range dir {
		if i-1 == input {
			// fmt.Println(f.Name())
			// fmt.Println(homeDir + "/Virtual Machines.localized" + "/" + f.Name())
			dir2, err := ioutil.ReadDir(homeDir + "/Virtual Machines.localized" + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}

			for _, d := range dir2 {
				if filepath.Ext(d.Name()) == ".vmx" {
					// fmt.Println(homeDir + "/Virtual Machines.localized" + "/" + f.Name() + "/" + d.Name())

					fullPath := homeDir + "/Virtual Machines.localized" + "/" + f.Name()
					vmxImage := d.Name()

					vmrunPath, _ := exec.LookPath("vmrun")
					cmdRun := &exec.Cmd{
						Path:   vmrunPath,
						Args:   []string{vmrunPath, "-T", "fusion", "start", vmxImage, "nogui"},
						Dir:    fullPath,
						Stdout: os.Stdout,
						Stderr: os.Stderr,
					}

					// fmt.Println( cmdRun.String() )

					if err := cmdRun.Run(); err != nil {
						fmt.Println("Error:", err)
					}
				}
				// 	fmt.Println(d.Name())
			}
			// fmt.Println(dir2)
		}
	}
}
