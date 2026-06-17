package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateFloat(inputNumber string) (int, error) {

	inputNumber = strings.TrimSpace(inputNumber)
	inputNumber = strings.ReplaceAll(inputNumber, " ", "")
	inputNumber = strings.ReplaceAll(inputNumber, "\t", "")

	result, err := strconv.ParseFloat(inputNumber, 64)
	if err != nil {
		return int(result), fmt.Errorf("ERROR : Enter number only")
	}

	return int(result), nil

}