package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Maths interface {
	Add() int
	Sub() int
}

type Value struct {
	InputOne int
	InputTwo int
	Result   int
}

func (a *Value) Add() int {

	a.Result = a.InputOne + a.InputTwo
	return a.Result
}
func (s *Value) Sub() int {

	s.Result = s.InputOne - s.InputTwo
	return s.Result
}
func isValid(number string) (int, error) {

	number = strings.TrimSpace(number)
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "\t", "")

	result, err := strconv.Atoi(number)
	if err != nil {
		return result, fmt.Errorf("ERROR : Enter number only")
	}

	return result, nil

}

func typeConversion(input any) (int, error) {
	
	switch v := input.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("Invalid Input type")
	}
}

func main() {

	var calculate Maths
	var inputValueOne, inputValueTwo any = 6, 5

	inputOne, err := typeConversion(inputValueOne)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputTwo, err := typeConversion(inputValueTwo)
	if err != nil {
		fmt.Println(err)
		return
	}

	calculate = &Value{
		InputOne: inputOne,
		InputTwo: inputTwo,
	}

	fmt.Println("Enter the choice \n1) ADD\n2) SUB")
	scanner.Scan()
	inputChoice := scanner.Text()
	choice, err := isValid(inputChoice)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch choice {
	case 1:
		fmt.Printf("\nThe sum of %d and %d is %d", inputValueOne, inputValueTwo, calculate.Add())
	case 2:
		fmt.Printf("\nThe difference between %d and %d is %d", inputValueOne, inputValueTwo, calculate.Sub())
	default:
		fmt.Println("Invalid Input, Enter number 1 or 2.")
	}

}
