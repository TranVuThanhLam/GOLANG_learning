package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type mapstring map[string]string

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(InputFile, OutputFile string) FileManager {
	return FileManager{
		InputFilePath:  InputFile,
		OutputFilePath: OutputFile,
	}
}

func (f FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(f.InputFilePath)

	if err != nil {
		return nil, errors.New("Could not open file")
	}

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Reading the file content failed!")
	}

	file.Close()
	return lines, nil
}

func (f FileManager) WriteResult(data any) error {
	file, err := os.Create(f.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file")
	}
	// w io.Writer = interface{}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data to json")
	}

	file.Close()
	return nil
}
