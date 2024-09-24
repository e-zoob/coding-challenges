package main

import (
	"ccwc/lib"
	"fmt"
	"os"
)

func main() {

	path := os.Args[1]

	size := lib.GetBytes(path)

	fmt.Printf("%d %s", size, path)

}
