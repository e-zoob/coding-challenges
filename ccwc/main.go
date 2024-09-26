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

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening the file", err)
	}

	defer file.Close()

	switch flag {
	case "-c":
		size := lib.CountBytes(file)
		fmt.Printf("%d %s", size, path)

	case "-l":
		lines := lib.CountLines(file)
		fmt.Printf("%d %s", lines, path)

	case "-w":
		words := lib.CountWords(file)
		fmt.Printf("%d %s", words, path)

	case "-m":
		chars := lib.CountChars(file)
		fmt.Printf("%d %s", chars, path)

	default:
		line, words, chars := lib.CountAll(file)

		fmt.Printf("%d, %d, %d, %s", line, words, chars, path)
	}

}
