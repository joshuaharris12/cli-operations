package main

import (
	"fmt"
	"go-ls/cmd"
	"os"
)

func main() {
	numArgs := len(os.Args)
	var directory string
	if numArgs > 2 {
		fmt.Printf("Expected 2 arguments, got %d\n", numArgs)
	}
	if numArgs == 1 {
		directory = "."
	} else {
		directory = os.Args[1]
	}

	output, err := cmd.Execute(directory)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(output)
}
