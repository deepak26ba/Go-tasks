package validation

import "fmt"

func IsValidString(input string) (string, error) {

	for i := 0; i < len(input); i++ {
		if !(input[i] >= 'A' && input[i] <= 'Z' || input[i] >= 'a' && input[i] <= 'z') {
			return "", fmt.Errorf("Name should not contain character other than Letters")
		}
	}
	return input, nil

}
