package main

import (
	"fmt"
	"go-cat/cmd"
	"os"
)

func main() {
	numArgs := len(os.Args)
	if numArgs == 1 {
		fmt.Printf("Error: Expected 1 argument, got %d\n", numArgs-1)
		os.Exit(1)
	}
	filePath := os.Args[1]
	err := cmd.Execute(filePath)
	if err != nil {
		fmt.Printf("Error: %w\n", err)
		os.Exit(1)
	}
}
