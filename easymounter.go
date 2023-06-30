// EASY MOUNTER BY HHK02

package main

import (
	"fmt"
	"os"
)

var source string
var target string
var mountprocess string = "mount -t auto " + source + " " + target

func main() {
	fmt.Println("EASY MOUNTER by hhk02")
	fmt.Println("Write source image: ")
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
		fmt.Println("Write source image: ")
		fmt.Scanln(&source)
	} else {
		fmt.Println("Mounting: " + source + "to " + target)
		os.StartProcess(mountprocess, nil, nil)
		fmt.Println(("Done!"))
	}
}
