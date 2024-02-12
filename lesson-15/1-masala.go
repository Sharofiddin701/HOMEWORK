package main

import "fmt"

type Animal interface {
	MakeSound() string
	Move() string
}

func PrintSound(a Animal) {

	fmt.Printf("Sound %s\n", a.MakeSound())

	fmt.Printf("Movement %s\n", a.Move())
}

type Dog struct {
}

func (d Dog) MakeSound() string {
	return "Wolf"
}

func (d Dog) Move() string {
	return "Tunning"
}

type Bird struct{}

func (b Bird) MakeSound() string {
	return "Tweet"

}
func (b Bird) Move() string {
	return "Flying"

}

func main() {

	dog := Dog{}
	bird := Bird{}

	PrintSound(dog)
	PrintSound(bird)

}
