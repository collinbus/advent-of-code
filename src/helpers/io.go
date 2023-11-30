package helpers

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(path string) []string {
	file, err := os.Open(path) // specify the file name here
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	lines := make([]string, 0)

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Use scanner to read the file line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
