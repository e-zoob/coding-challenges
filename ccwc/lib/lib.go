package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func CountBytes(text string) int64 {
	return int64(len(text))
}

func CountLines(text string) int {
	var count int
	var target = []byte("\n")

	reader := strings.NewReader(text)
	buffer := make([]byte, 32*1024)

	for {
		read, err := reader.Read(buffer)
		if err != nil {
			break
		}
		count += bytes.Count(buffer[:read], target)
	}
	return count
}

func CountWords(text string) int {
	var count int
	delim := byte('\n')
	reader := bufio.NewReader(strings.NewReader(text))

	for {
		line, err := reader.ReadString(delim)
		if err != nil {
			break
		}

		count += len(strings.Fields(line))
	}
	return count
}

func CountChars(text string) int {
	var count int
	delim := byte('\n')
	reader := bufio.NewReader(strings.NewReader(text))

	for {
		line, err := reader.ReadString(delim)
		count += utf8.RuneCountInString(line)
		if err != nil {
			break
		}

	}
	return count
}

func CountAll(text string) (int, int, int) {
	lines := CountLines(text)
	words := CountWords(text)
	chars := CountChars(text)

	return lines, words, chars

}

func ProcessText(text string, flag string) map[string]int {
	result := map[string]int{}
	switch flag {
	case "-c":
		result["bytes"] = int(CountBytes(text))
	case "-l":
		result["lines"] = CountLines(text)
	case "-w":
		result["words"] = CountWords(text)
	case "-m":
		result["chars"] = CountChars(text)
	default:
		lines, words, chars := CountAll(text)
		result["lines"] = lines
		result["words"] = words
		result["chars"] = chars
	}
	return result
}
func PrintResult(result map[string]int, path string, flag string) {
	switch flag {
	case "-c":
		fmt.Printf("%d %s\n", result["bytes"], path)
	case "-l":
		fmt.Printf("%d %s\n", result["lines"], path)
	case "-w":
		fmt.Printf("%d %s\n", result["words"], path)
	case "-m":
		fmt.Printf("%d %s\n", result["chars"], path)
	default:
		fmt.Printf("%d %d %d %s\n", result["lines"], result["words"], result["chars"], path)
	}
}
