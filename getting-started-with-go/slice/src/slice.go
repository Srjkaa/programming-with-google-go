package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getUserInput() string {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println(err)
	}
	return userInput
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
		enteredValue := getUserInput()
		if enteredValue[0] == 'X' {
			return
		}
		enteredNumber, _ := strconv.Atoi(enteredValue)
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
