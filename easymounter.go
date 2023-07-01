// EASY MOUNTER BY HHK02

package main

import (
	"fmt"
	"os/exec"
)

var args []string
var mount string = "mount"
var source string
var target string

func main() {

	fmt.Println("EASY MOUNTER by hhk02")
	fmt.Println("Write source path: ")
	fmt.Scanln(&source)
	fmt.Println("Write target path: ")
	fmt.Scanln(&target)

	if target == "" {
		fmt.Println("Write target path: ")
		fmt.Scanln(&target)
	} else {
		fmt.Println("Target path: " + target)
	}

	if source == "" {
		fmt.Println("Invalid value!")
		fmt.Println("Write source path: ")
		fmt.Scanln(&source)
	} else {

		fmt.Println("Mounting: " + source + " to " + target)
		cmd := exec.Command(mount, source, target)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("Error")
		}

		fmt.Println(string(out))
		fmt.Println(("Done!"))
	}
}
