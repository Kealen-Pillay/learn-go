package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// read contents of text file - from cmd line
	// open file
	// io.Copy
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error file doesnt exist")
		os.Exit(-1)
	} 
	io.Copy(os.Stdout, file)
	
}