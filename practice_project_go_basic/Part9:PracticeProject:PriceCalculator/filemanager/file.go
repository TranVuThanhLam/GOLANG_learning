package filemanager

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FileName = "prices.txt"

func ReadFileFloatData() []float64 {
	fileDataByte, err := os.ReadFile(FileName)
	if err != nil {
		fmt.Println("Fails to read file")
	}

	fileDataString := string(fileDataByte)

	fileDataStringSlice := strings.Split(fileDataString, "\n")

	var fileDataFloatSlice []float64
	for _, dataString := range fileDataStringSlice {
		dataFloat, err := strconv.ParseFloat(dataString, 64)
		if err != nil {
			fmt.Println("Data in file is note float")
		}
		fileDataFloatSlice = append(fileDataFloatSlice, dataFloat)
	}

	return fileDataFloatSlice
}
