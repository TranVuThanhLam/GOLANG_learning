package convertion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := []float64{}
	for _, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}
		floats = append(floats, floatValue)
	}
	return floats, nil
}
