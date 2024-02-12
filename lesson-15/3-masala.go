package main

import (
	"fmt"
)

type Book interface {
	Autor() string
	Title() string
}



type Autor struct {
	autor string
}

func (c Autor) Autor() string {
	return c.Autor()
}

type Title struct {
	title string
}

func (r Title) Title() string {
	return r.Title()
}

func pshd(s Book) {

	fmt.Printf("Kitob: %v\n", s.Autor(),s.Title())

}
func main() {

	r := Autor{"Muhammad Yusuf"}
	pshd(r)
	c := Title{"Sherlar to'plami"}
	pshd(c)

}
