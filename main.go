package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func saveFilesAndContents(directory string) {
	outputFile := "output.txt"

	info, err := os.Stat(directory)
	if err != nil || !info.IsDir() {
		fmt.Printf("Error: The directory '%s' does not exist.\n", directory)
		return
	}

	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening output file: %v\n", err)
		return
	}
	defer f.Close()

	header := fmt.Sprintf("\n\n=== Checking files in directory: %s ===\n\n", directory)
	if _, err := f.WriteString(header); err != nil {
		fmt.Printf("Error writing header: %v\n", err)
		return
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			separator := strings.Repeat("=", 50)
			f.WriteString(fmt.Sprintf("File: %s\n%s\n", filename, separator))

			filePath := filepath.Join(directory, filename)
			content, err := os.ReadFile(filePath)
			if err != nil {
				f.WriteString(fmt.Sprintf("Could not read file: %v\n", err))
			} else {
				f.WriteString(string(content) + "\n")
			}
			f.WriteString("\n" + strings.Repeat("-", 50) + "\n\n")
		}
	}

	fmt.Printf("File contents appended to %s\n", outputFile)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dir-content-saver <directory_path>")
		os.Exit(1)
	}

	directory := os.Args[1]
	saveFilesAndContents(directory)
}
