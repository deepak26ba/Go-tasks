package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for {

		fmt.Print("\nEnter the choice\n1) Initailize Array \n2) Return Position \n3) Double Value \n4) Sort\n5) EXIT")
		fmt.Print("\nEnter the number : ")
		scanner.Scan()
		number := scanner.Text()

		value, err := isValid(number)
		if err != nil {
			fmt.Println(err)
		}

		switch value {
		case 1:
			initailizeArray()
		case 2:
			returnPosition()
		case 3:
			doubleValue()
		case 4:
			sortArray()
		case 5:
			fmt.Println("Loop exited. Program finished.")
			return

		default:
			fmt.Println("Invalid Input, Enter number 1 to 5")
		}

	}

}

func initailizeArray() {

	numberArray := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nameArray := []string{"GO", "Python", "PHP", "React", ".Net"}
	fmt.Println("\n", numberArray, "\n", nameArray)
}

func returnPosition() {

	numberArray := [8]int{1, 2, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array Elments : ", numberArray)

	fmt.Print("\nEnter the target number : ")
	scanner.Scan()
	target := scanner.Text()

	targetValue, err := isValid(target)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := checkPosition(numberArray, targetValue)
	fmt.Println(result)

}

func doubleValue() {

	var doubleValueArray [7]int

	fmt.Print("\nEnter the first array elment : ")
	scanner.Scan()
	number := scanner.Text()

	value, err := isValid(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	doubleValueArray[0] = value
	fmt.Printf("array_nums[0] = %d \n", value)

	for i := 1; i < len(doubleValueArray); i++ {
		doubleValueArray[i] = doubleValueArray[i-1] * 2
		fmt.Printf("array_nums[%d] = %d \n", i, doubleValueArray[i])
	}

}

func sortArray() {

	sortedArray := [6]int{8, 3, 2, 1, 3, 2}
	fmt.Println("Array : ", sortedArray)
	for i := range len(sortedArray) {
		for j := i + 1; j < len(sortedArray); j++ {
			if sortedArray[i] > sortedArray[j] {
				sortedArray[i], sortedArray[j] = sortedArray[j], sortedArray[i]
			} else if sortedArray[i] < sortedArray[j] || sortedArray[i] == sortedArray[j] {
				continue
			}

		}

	}

	fmt.Print("Ascending order : ")
	for i := range len(sortedArray) {
		fmt.Print(sortedArray[i], " ")

	}

	fmt.Print("\nDecending order : ")
	for i := len(sortedArray) - 1; i >= 0; i-- {
		fmt.Print(sortedArray[i], " ")

	}

}

func checkPosition(checkArray [8]int, target int) string {

	fisrtOccurence := true

	for i := range len(checkArray) {
		for j := 1; j < len(checkArray)-i-1; j++ {
			if checkArray[i]+checkArray[j] == target {
				if checkArray[i] == checkArray[j] {
					return ""
				}

				if fisrtOccurence {
					fmt.Printf("shortest path : %d,%d - %d,%d\n", i, j, checkArray[i], checkArray[j])
					fmt.Println("other possibilities : ")
					fisrtOccurence = false
				} else {
					fmt.Printf("%d,%d - %d,%d\n", i, j, checkArray[i], checkArray[j])

				}

			}

		}

	}

	if fisrtOccurence == true {
		fmt.Printf("\nEntered number %d is not possible to find sum in the array \n", target)
	}
	return ""

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
