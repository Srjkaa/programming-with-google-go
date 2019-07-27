package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

type AnimalType string
const (
	Cow AnimalType = "cow"
	Bird = "bird"
	Snake = "snake"
)

func (animal *Animal) InitializeAnimal(animalType AnimalType) *Animal {
	switch animalType {
	case Cow:
		animal.food = "grass"
		animal.locomotion = "walk"
		animal.noise = "moo"
	case Bird:
		animal.food = "worms"
		animal.locomotion = "fly"
		animal.noise = "peep"
	case Snake:
		animal.food = "mice"
		animal.locomotion = "slither"
		animal.noise = "hsss"
	}
	return animal
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func (animal *Animal) callAppropriateMethod (enteredCommand string) (string, error) {
	switch enteredCommand {
	case "eat":
		return animal.Eat(), nil
	case "move":
		return animal.Move(), nil
	case "speak":
		return animal.Speak(), nil
	}
	return "", errors.New("unknown command")
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	scanner.Scan()
	return scanner.Text()
}

func getAnimalByType(enteredAnimal AnimalType, cow, bird, snake *Animal) (*Animal, error) {
	switch enteredAnimal {
	case Cow:
		return cow, nil
	case Bird:
		return bird, nil
	case Snake:
		return snake, nil
	}
	return &Animal{}, errors.New("unknown animal")
}

func main() {
	var cow Animal
	cow = *cow.InitializeAnimal(Cow)
	var bird Animal
	bird = *bird.InitializeAnimal(Bird)
	var snake Animal
	snake = *snake.InitializeAnimal(Snake)

	for {
		userInputAsList := strings.Split(getUserInput(), " ")
		enteredAnimal := AnimalType(userInputAsList[0])
		enteredCommand := userInputAsList[1]
		animal, err := getAnimalByType(enteredAnimal, &cow, &bird, &snake)
		if err != nil {
			fmt.Println(err)
			continue
		}

		commandResult, err := animal.callAppropriateMethod(enteredCommand)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(commandResult)
	}
}