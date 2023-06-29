// EASY MOUNTER BY HHK02

package main

import (
	"context"
	"fmt"

	"github.com/dell/gofsutil"
)

var source string
var target string = "/mnt/target"

func main() {
	fmt.Println("EASY MOUNTER by hhk02")
	fmt.Println("Write source image: ")
	fmt.Scanln(&source)

	if source == "" {
		fmt.Println("Invalid value!")
		fmt.Println("Write source image: ")
		fmt.Scanln(&source)
	} else {
		fmt.Println("Mounting: " + source)
		gofsutil.Mount(context.Background(), source, target, "ext4", "rw")
		fmt.Println(("Done!"))
	}
}
