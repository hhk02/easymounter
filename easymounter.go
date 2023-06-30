// EASY MOUNTER BY HHK02

package main

import (
	"context"
	"fmt"

	"github.com/dell/gofsutil"
)

var source string
var target string

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
		gofsutil.Mount(context.Background(), source, target, "auto", "rw,relatime,errors=remount-ro")
		fmt.Println(("Done!"))
	}
}
