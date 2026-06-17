package calculate_length

import (
	"fmt"
)

func CalculateLength() (*string, error) {

	var inputString string

	fmt.Print("\nEnter a string : ")
	_, err := fmt.Scanf("%s\n", &inputString)
	if err != nil {
		return nil, fmt.Errorf("Error : %v", err)
	}

	pointer := &inputString

	return pointer, nil
}
