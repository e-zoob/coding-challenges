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
		size := lib.CountBytes(path)
		fmt.Printf("%d %s", size, path)

	case "-l":
		lines := lib.CountLines(path)
		fmt.Printf("%d %s", lines, path)

	case "-w":
		words := lib.CountWords(path)
		fmt.Printf("%d %s", words, path)
	}

}
