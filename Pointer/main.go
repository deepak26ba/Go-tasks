package main

import (
	"bufio"
	"fmt"
	"os"
	"task12/add_numbers"
	"task12/calculate_length"
	"task12/sort_array"
	"task12/struct_using_ptr"
	"task12/swap_elements"
	"task12/validation"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for {

		fmt.Print("\nEnter the choice\n1) Calculate the length of the string \n2) Sort an array\n3) Add numbers\n4) Employee Details\n5) Swap Elements\n6) EXIT")

		fmt.Print("\nEnter the number : ")
		scanner.Scan()
		choiceNumber := scanner.Text()
		value, err := validation.ValidateNumber(choiceNumber)
		if err != nil {
			fmt.Println(err)
		}

		switch value {
		case 1:
			inputString, err := calculate_length.CalculateLength()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("The length of the given string %v is : %d\n", *inputString, len(*inputString))

		case 2:

			inputSlice, err := sort_array.SortArray()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The elements in the array after sorting \nAscending Order")
			for i := range len(inputSlice) {

				fmt.Printf("element - %d : %d \n", i+1, *inputSlice[i])
			}
			fmt.Println("\nDescending Order")
			for i, j := len(inputSlice)-1, 1; i >= 0; i-- {
				fmt.Printf("element - %d : %d \n", j, *inputSlice[i])
				j++
			}

		case 3:

			firstNumber, secondNumber, sum, err := add_numbers.AddNumbers()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("\nThe sum of %d and %d  is %d\n", firstNumber, secondNumber, sum)

		case 4:

			employeePTR, err := struct_using_ptr.StructUsingPtr()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("\n\tEmployee Details ")
			fmt.Println("Employee ID     : ", employeePTR.ID)
			fmt.Println("Employee Name   : ", employeePTR.Name)
			fmt.Println("Employee DOB    : ", employeePTR.DateOfBirth)

		case 5:
			
			inputValueOne, inputValueTwo, inputValueThree, err := swap_elements.SwapElements()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("\nThe value after swapping are : ")
			fmt.Printf("element 1 = %d \nelement 2 = %d \nelement 3 = %d\n", inputValueOne, inputValueTwo, inputValueThree)

		case 6:
			fmt.Println("Loop exited. Program finished.")
			return

		default:
			fmt.Println("Invalid Input, Enter number 1 to 6")
		}
	}
}
