package lib

import (
	"fmt"
	"os"
)

func GetBytes(path string) int64 {
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
