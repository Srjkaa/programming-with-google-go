package main

import (
	"fmt"
	"regexp"
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

func main() {
	numbers := make([]int, 3)
	re := regexp.MustCompile(`[a-zA-Z]`)
	index := 0
	for {
		fmt.Print("Enter 'X' for quit or an integer: ")
		enteredValue := getUserInput()
		if enteredValue[0] == 'X' {
			fmt.Println(numbers)
			return
		}
		if re.Match([]byte(enteredValue)) {
			continue
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
		index++
	}
}
