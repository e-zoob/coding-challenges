package main

import (
	"ccwc/lib"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Println("USAGE: <program> <flag> <path> OR stdin | <program> <flag> ")
	}
	if len(os.Args) == 1 {
	}

	var flag, path string
	flags := []string{"-c", "-l", "-m", "-w"}

	if len(os.Args) == 2 && !slices.Contains(flags, os.Args[1]) {
		path = os.Args[1]
	} else if len(os.Args) == 2 && slices.Contains(flags, os.Args[1]) {
		readString, err := io.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		result := lib.ProcessText(string(readString), os.Args[1])
		lib.PrintResult(result, "", os.Args[1])
		return

	} else {
		flag = os.Args[1]
		path = os.Args[2]
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	const maxSize int64 = 10 * 1024 * 1024

	if fileInfo.Size() > maxSize {
		fmt.Println("File too large.")
		return
	} else {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening the file", err)
			return
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading the file", err)
			return
		}

		text := string(content)
		result := lib.ProcessText(text, flag)
		lib.PrintResult(result, path, flag)
	}

}
