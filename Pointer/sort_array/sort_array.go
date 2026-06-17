package sort_array

import (
	"bufio"
	"fmt"
	"os"
	"task12/validation"
)

var scanner = bufio.NewScanner(os.Stdin)

func SortArray() ([]*int, error) {

	var emptySlice []*int

	fmt.Print("\nEnter the number of elements to store in the array : ")
	scanner.Scan()
	sliceSize := scanner.Text()
	size, err := validation.ValidateNumber(sliceSize)
	if err != nil {
		return nil, err
	}

	if size < 0 {
		return nil, fmt.Errorf("Array length can't be negative")
	}

	fmt.Printf("Input %d number of elements in the array : \n", size)
	inputSlice := make([]*int, size)

	for i := range size {

		scanner.Scan()
		element := scanner.Text()
		value, err := validation.ValidateNumber(element)
		if err != nil {

			return emptySlice, err
		}
		inputSlice[i] = &value

	}
	
	sortSlice(inputSlice)

	return inputSlice, nil

}

func sortSlice(inputArray []*int) {

	for i := range len(inputArray) {
		for j := i + 1; j < len(inputArray); j++ {
			if *inputArray[i] > *inputArray[j] {
				inputArray[i], inputArray[j] = inputArray[j], inputArray[i]
			}
		}
	}
}
