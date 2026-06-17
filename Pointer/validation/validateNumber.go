package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateNumber(inputNumber string) (int, error) {

	inputNumber = strings.TrimSpace(inputNumber)
	inputNumber = strings.ReplaceAll(inputNumber, " ", "")
	inputNumber = strings.ReplaceAll(inputNumber, "\t", "")

	result, err := strconv.Atoi(inputNumber)
	if err != nil {
		return result, fmt.Errorf("ERROR : Enter number only")
	}

	return result, nil

}
