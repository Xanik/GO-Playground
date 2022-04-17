package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name is missing")
		os.Exit(1)
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	// Copy the content from file's Reader into Stdout's Writer
	io.Copy(os.Stdout, file)
}
