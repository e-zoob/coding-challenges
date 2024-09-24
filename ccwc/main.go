package main

import (
	"ccwc/lib"
	"fmt"
	"os"
)

func main() {

	path := os.Args[2]

	switch flag := os.Args[1]; flag {
	case "-c":
		size := lib.GetBytes(path)
		fmt.Printf("%d %s", size, path)

	}

}
