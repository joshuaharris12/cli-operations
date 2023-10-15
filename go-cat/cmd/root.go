package cmd

import "os"

func Execute(filePath string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(fileBytes)
	if err != nil {
		return err
	}
	os.Stdout.WriteString("\n")
	return nil
}
