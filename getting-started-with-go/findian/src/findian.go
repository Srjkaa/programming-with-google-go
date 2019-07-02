package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var givenString string
	fmt.Println("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		givenString += scanner.Text()
	}

	givenStringInLowerCase := strings.ToLower(givenString)
	isStartsWithRuneI := strings.IndexRune(givenStringInLowerCase, 'i') == 0
	isEndsWithRuneN := strings.IndexRune(givenStringInLowerCase, 'n') == len(givenString) - 1
	isContainsRuneA := strings.ContainsRune(givenStringInLowerCase, 'a')
	if (isStartsWithRuneI && isEndsWithRuneN && isContainsRuneA) {
		fmt.Println("Found!")
		return
	}
	fmt.Println("Not Found!")
}
