package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueStr := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueStr), 0644) // 0644 is the file permission (read/write for owner, read for others)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil { // nil stands for an abscense of a useful value, in this case it means that there was an error reading the file.
		return 1000, errors.New("Failed to find file")
	}
	valueText := string(data) // Convert the byte slice to a string. The byte slice can't be converted directly to a float, so we first convert it to a string and then parse it as a float.
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}
	return value, nil
}
