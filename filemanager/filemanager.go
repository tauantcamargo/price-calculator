package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error opening file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Error reading file")
	}

	file.Close()

	return lines, nil
}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("Error creating file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Error writing JSON")
	}

	file.Close()

	return nil
}
