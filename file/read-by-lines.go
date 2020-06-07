package main

import (
	"bufio"
	"os"
)

func ReadFileToStringsSlice(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return lines, nil
}
