package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(list []int, index int) {
	temporary := list[index]
	list[index] = list[index+1]
	list[index+1] = temporary
}

func bubbleSort(list []int) {
	length := len(list)
	var isSorted bool

	for range list {
		isSorted = true
		for index, _ := range list {
			if index == length-1 {
				break
			}
			if list[index] > list[index+1] {
				isSorted = false
				swap(list, index)
			}
		}
		if isSorted {
			break
		}
	}
}

func main() {
	listByUserInput := make([]int, 0, 10)
	fmt.Println("Enter numbers in a sequence up to 10 integers:")
	scanner := bufio.NewScanner(os.Stdin)
	var userInput []string
	if scanner.Scan() {
		userInput = strings.SplitN(strings.Trim(scanner.Text(), " "), " ", 10)
	}
	for _, value := range userInput {
		number, _ := strconv.Atoi(value)
		listByUserInput = append(listByUserInput, number)
	}
	bubbleSort(listByUserInput)
	fmt.Println(listByUserInput)
}
