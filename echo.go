// echo.go
// Simple echo utility
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	buffer := " "
	if len(os.Args) > 1 {
		buffer = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(buffer)
}
