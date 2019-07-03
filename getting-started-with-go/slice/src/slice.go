package main

import (
	"fmt"
	"sort"
)

// TODO: implement quit when 'X' is entered
func getUserEnteredValue() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err)
	}
	return number
}

func replaceInitialSliceValuesWithGivenNumber(slice []int, number int) []int {
	index := sort.SearchInts(slice, 0)
	slice[index] = number
	return slice
}

// TODO: implement case when user wants to input more than 3 values but with initial slice length equals to 3
func main() {
	numbers := make([]int, 3)
	for index := range numbers {
		fmt.Print("Enter 'X' for quit or an integer: ")
		enteredNumber := getUserEnteredValue()
		switch {
		case index < 3:
			replaceInitialSliceValuesWithGivenNumber(numbers, enteredNumber)
		default:
			numbers = append(numbers, enteredNumber)
		}
		sort.Ints(numbers)
		fmt.Println(numbers)
	}
}
