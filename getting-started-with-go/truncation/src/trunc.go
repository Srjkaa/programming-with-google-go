package main

import "fmt"

func main() {
	fmt.Printf("Enter a floating point number: \n")
	var givenNumber float64
	_, err := fmt.Scan("%f", &givenNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Truncated given number is: %d \n", int64(givenNumber))
}
