package fileutils

import (
	"bufio"
	"os"
)

func ReadFile(path string) (string, error) {
	var fileContent string
	file, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fileContent, nil
}
