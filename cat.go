// cat.go
// Read file given argument
// Wrote this some time ago, might need some cleanup
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(f *os.File) {

	if f == nil {
		return
	}

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')

	// While there is no error, read lines
	for err == nil {
		fmt.Print(line)
		line, err = reader.ReadString('\n')
	}

	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("You must provide an argument")
		return
	}

	file, err := os.Open(args[1])

	// If there is an error opening the file
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	cat(file)
}
