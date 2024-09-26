package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func CountBytes(path string) int64 {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening the file", err)
		return 0
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println("Error getting file info:", err)
		return 0
	}

	return fileInfo.Size()

}

func CountLines(path string) int {
	var count int
	var read int
	var target []byte = []byte("\n")

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening the file", err)
		return 0
	}
	defer file.Close()

	buffer := make([]byte, 32*1024)
	reader := bufio.NewReader(file)

	for {
		read, err = reader.Read(buffer)
		if err != nil {
			break
		}
		count += bytes.Count(buffer[:read], target)
	}
	if err == io.EOF {
		return count
	}
	return count
}

func CountWords(path string) int {
	var count int
	delim := byte('\n')

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening the file", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(delim)
		if err != nil {
			break
		}

		if err == io.EOF {
			return count
		}
		count += len(strings.Fields(line))
	}
	return count
}
