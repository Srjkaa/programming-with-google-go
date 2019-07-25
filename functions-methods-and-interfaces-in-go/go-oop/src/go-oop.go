package main

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.noise
}

func main() {

}