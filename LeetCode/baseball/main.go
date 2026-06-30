package main

import "strconv"

func callpoint(operations []string) int {

	pointslice := make([]int, 0)

	for i := range operations {

		if operations[i] == "C" {
			pointslice = pointslice[:len(pointslice)-1]
		} else if operations[i] == "D" {
			num := pointslice[len(pointslice)-1]
			pointslice = append(pointslice, num*2)
		} else if operations[i] == "+" {
			num1 := pointslice[len(pointslice)-1]
			num2 := pointslice[len(pointslice)-2]
			pointslice = append(pointslice, num1+num2)
		} else {
			num, _ := strconv.Atoi(operations[i])
			pointslice = append(pointslice, num)
		}
	}

	score := 0
	for i := 0; i < len(pointslice); i++ {
		score += pointslice[i]
	}

	return score
}
