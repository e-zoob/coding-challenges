package main

import (
	"ccwc/lib"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: <program> <flag> <path>")
	}

	var flag, path string

	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		flag = os.Args[1]
		path = os.Args[2]
	}

	switch flag {
	case "-c":
		size := lib.CountBytes(path)
		fmt.Printf("%d %s", size, path)

	case "-l":
		lines := lib.CountLines(path)
		fmt.Printf("%d %s", lines, path)

	case "-w":
		words := lib.CountWords(path)
		fmt.Printf("%d %s", words, path)

	case "-m":
		chars := lib.CountChars(path)
		fmt.Printf("%d %s", chars, path)

	default:
		line, words, chars := lib.CountAll(path)

		fmt.Printf("%d, %d, %d, %s", line, words, chars, path)
	}

}
