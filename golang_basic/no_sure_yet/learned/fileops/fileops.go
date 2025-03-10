package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to find balance.")
	}
	ValueText := string(data)
	value, err := strconv.ParseFloat(ValueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}
	return value, nil
}

func WriteFloatToFile(Value float64, fileName string) {
	ValueText := fmt.Sprint(Value)
	os.WriteFile(fileName, []byte(ValueText), 0644)
}
