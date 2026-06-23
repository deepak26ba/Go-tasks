package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var count int
	primeNumber:=make( chan int)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number : ")
	scanner.Scan()
	number := scanner.Text()

	number = strings.TrimSpace(number)
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "\t", "")

	value, err := checkNumber(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	go prime(value,primeNumber)

	fmt.Println("The prime numbers are ")
	
	for i := range primeNumber{
		fmt.Println(i)
		count ++
	}
	fmt.Printf("Total prime number from 1 to %d is %d",value,count)

	
}

func prime(value int, primeNumber chan int) {


	for i := 1; i <= value; i++ {
		if checkPrime(i) {
			primeNumber <- i
		}
	}
	close(primeNumber)

}

func checkPrime(vaule int) bool {
	if vaule <= 1 {
		return false
	}

	for i := 2; i*i <= vaule; i++ {
		if vaule%i == 0 {
			return false
		}
	}
	return true

}

func checkNumber(value string) (int, error) {

	result, err := strconv.Atoi(value)
	if err != nil {
		return result, fmt.Errorf("ERROR : Enter number only")
	}
	return result, nil

}
