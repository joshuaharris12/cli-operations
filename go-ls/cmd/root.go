package cmd

import (
	"fmt"
	"os"
)

func Execute(directory string) (string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", err
	}
	output := ""
	for _, file := range files {
		output += fmt.Sprintf("%s\t", file.Name())
	}
	if len(output) != 0 {
		output += "\n"
	}
	return output, nil
}
