package swap_elements

import (
	"bufio"
	"fmt"
	"os"
	"task12/validation"
)

var scanner = bufio.NewScanner(os.Stdin)

func SwapElements() (int, int, int, error) {

	var inputValueOne, inputValueTwo, inputValueThree int
	var err error

	fmt.Print("\nEnter the value of 1st element : ")
	scanner.Scan()
	inputOne := scanner.Text()
	inputValueOne, err = validation.ValidateFloat(inputOne)
	if err != nil {
		return inputValueOne, inputValueTwo, inputValueThree, err
	}

	fmt.Print("Enter the value of 2nd element : ")
	scanner.Scan()
	inputTwo := scanner.Text()
	inputValueTwo, err = validation.ValidateFloat(inputTwo)
	if err != nil {
		return inputValueOne, inputValueTwo, inputValueThree, err
	}

	fmt.Print("Enter the value of 3rd element : ")
	scanner.Scan()
	inputThree := scanner.Text()
	inputValueThree, err = validation.ValidateFloat(inputThree)
	if err != nil {
		return inputValueOne, inputValueTwo, inputValueThree, err
	}

	fmt.Println("\nThe value before swapping are : ")
	fmt.Printf("element 1 = %d \nelement 2 = %d \nelement 3 = %d\n", inputValueOne, inputValueTwo, inputValueThree)

	inputValueOne, inputValueTwo, inputValueThree = inputValueThree, inputValueOne, inputValueTwo

	return inputValueOne, inputValueTwo, inputValueThree, nil
}
