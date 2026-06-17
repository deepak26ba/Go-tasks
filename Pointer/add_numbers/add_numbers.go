package add_numbers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task12/validation"
)

var scanner = bufio.NewScanner(os.Stdin)

func AddNumbers() (int, int, int, error) {

	var sum, inputNumberOne, inputNumberTwo int
	var err error

	fmt.Print("\nEnter the first number : ")
	scanner.Scan()
	inputOne := scanner.Text()
	inputNumberOne, err = validation.ValidateNumber(inputOne)
	if err != nil {
		return inputNumberOne, inputNumberTwo, sum, err
	}

	fmt.Print("Enter the second number : ")
	scanner.Scan()
	inputTwo := scanner.Text()
	inputNumberTwo, err = validation.ValidateNumber(inputTwo)
	if err != nil {
		fmt.Println(err)
		return inputNumberOne, inputNumberTwo, sum, err
	}

	findSum(&inputNumberOne, &inputNumberTwo, &sum)

	return inputNumberOne, inputNumberTwo, sum, nil
}

func findSum(firstNumber, secondNumber, sum *int) {

	*sum = *firstNumber + *secondNumber
}

func isValid(inputNumber string) (int, error) {

	inputNumber = strings.TrimSpace(inputNumber)
	inputNumber = strings.ReplaceAll(inputNumber, " ", "")
	inputNumber = strings.ReplaceAll(inputNumber, "\t", "")

	result, err := strconv.Atoi(inputNumber)
	if err != nil {
		return result, fmt.Errorf("ERROR : Enter number only")
	}

	return result, nil

}
