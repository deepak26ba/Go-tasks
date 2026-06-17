package struct_using_ptr

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task12/validation"
	"time"
)

const layout = "02-01-2006"

var scanner = bufio.NewScanner(os.Stdin)

type Employee struct {
	ID          string
	Name        string
	DateOfBirth string
}

func StructUsingPtr() (*Employee, error) {

	var employeeDetail Employee
	//Getting the user name

	fmt.Print("\nEnter the Name : ")
	scanner.Scan()
	input := scanner.Text()
	inputName := strings.Fields(input)
	username := strings.Join(inputName, " ")
	name, err := validation.IsValidString(username)
	if err != nil {
		return nil, err
	}

	//getting the user dob

	date, err := getBirthDate()
	if err != nil {
		return nil, err
	}
	dob := date.Format(layout)

	//getting the user id

	fmt.Print("Enter the ID : ")
	scanner.Scan()
	inputID := scanner.Text()
	inputID = strings.TrimSpace(inputID)
	inputID = strings.ReplaceAll(inputID, " ", "")
	inputID = strings.ReplaceAll(inputID, "\t", "")

	employeeDetail = Employee{
		Name:        name,
		DateOfBirth: dob,
		ID:          inputID,
	}

	//creating a pointer variable
	employeePTR := &employeeDetail

	return employeePTR, nil
}

func getBirthDate() (time.Time, error) {

	var ageLimitBelow = time.Now().Year() - 50
	var ageLimitAbove = time.Now().Year() - 18
	var date string
	var defaultTime time.Time

	//Getting the input from user
	fmt.Print("Enter a your brithday date (DD-MM-YYYY) : ")
	fmt.Scanln(&date)
	resultdate, err := time.Parse(layout, date)
	if err != nil {
		return defaultTime, fmt.Errorf("Incorrect Date %v", err.Error())
	}

	//validating the age limit
	if ageLimitAbove < resultdate.Year() {
		return defaultTime, fmt.Errorf("Invalid Date : Your age is below the limit ")

	}
	if resultdate.Year() < ageLimitBelow {
		return defaultTime, fmt.Errorf("Invalid Date : Your age is above the limit  ")
	}

	return resultdate, nil
}
