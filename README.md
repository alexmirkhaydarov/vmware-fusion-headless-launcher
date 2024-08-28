# Running a VMWare Fusion VM Headless (no GUI)

A little handy utility to launch VMware fusion images in a headless mode. It does by listing all the existing `.vmx` images in the default directory and be able to choose which image to launch. Currently the existing vmrun tool doesn't list the path/directory where your images are stored, rather you have to provide the full path to launch the images in a headless mode. It saves me time from giving the full path to the `.vmx` image everytime I want to launch a vm in a headless mode.

- *This project is still in progress.*
- *Tested only on macos running on M1/M2 chip*
- *Tested with VMware Fusion 13.5.0 installed, and using the default `$HOME/Virtual Machines` directory for storing the `.vmx` images*

___

## Build
```bash
env GOOS=darwin GOARCH=arm64 go build -o vmruncli
```

## Add to path
```bash
mv vmruncli /usr/local/bin
```
