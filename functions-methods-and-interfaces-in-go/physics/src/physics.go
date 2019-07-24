package main

import (
	"fmt"
	"math"
)

func genDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func (float64) float64 {
	fn := func(time float64) float64 {
		return 0.5 * acceleration * math.Pow(time, 2) + initialVelocity * time + initialDisplacement
	}
	return fn
}

func promptUserInput(greeting string) float64 {
	fmt.Printf("%s: ", greeting)
	var value float64
	fmt.Scan(&value)
	return value
}

func main() {
	acceleration := promptUserInput("acceleration")
	initialVelocity := promptUserInput("initial velocity")
	initialDisplacement := promptUserInput("initial displacement")
	time := promptUserInput("time")
	fn := genDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Displacement after the entered time: %f\n", fn(time))
}
