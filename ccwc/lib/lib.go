package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func CountBytes(file *os.File) int64 {

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println("Error getting file info:", err)
		return 0
	}

	return fileInfo.Size()

}

func CountLines(file *os.File) int {
	var count int
	var err error
	var target []byte = []byte("\n")

	file.Seek(0, io.SeekStart)

	buffer := make([]byte, 32*1024)
	reader := bufio.NewReader(file)

	for {
		read, err := reader.Read(buffer)
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

func CountWords(file *os.File) int {
	var count int
	var err error
	delim := byte('\n')
	file.Seek(0, io.SeekStart)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(delim)
		if err != nil {
			break
		}

		count += len(strings.Fields(line))
	}
	if err == io.EOF {
		return count
	}
	return count
}

func CountChars(file *os.File) int {
	var count int
	var err error
	delim := byte('\n')
	file.Seek(0, io.SeekStart)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(delim)
		if err != nil {
			break
		}

		count += utf8.RuneCountInString(line)
	}
	if err == io.EOF {
		return count
	}
	return count
}

func CountAll(file *os.File) (int, int, int) {
	lines := CountLines(file)
	words := CountWords(file)
	chars := CountChars(file)

	return lines, words, chars

}
