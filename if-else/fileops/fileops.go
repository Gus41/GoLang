package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func SaveValueInFile(value float64, fileName string) {
	var valuteText = fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valuteText), 0644)
}
func ReadValueFromFile(fileName string) float64 {
	//if file not found, creates with default value 0.0
	data, err := os.ReadFile(fileName)
	if err != nil {
		SaveValueInFile(0, fileName)
		return 0.0
	}

	valueText := string(data)
	valueFloat, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0.0
	}

	return valueFloat
}
