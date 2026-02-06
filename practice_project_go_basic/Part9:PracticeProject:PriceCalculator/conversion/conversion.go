package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, StringValue := range strings {
		floatPrice, err := strconv.ParseFloat(StringValue, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, floatPrice)
	}

	return floats, nil
}
