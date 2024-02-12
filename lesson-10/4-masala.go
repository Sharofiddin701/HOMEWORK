package main

import "fmt"

type book struct {
	Nomi       string
	Aftori     string
	Mavjud     int
	Umumiysoni int
}

func (c *book) borrow(i int) {

	if c.Mavjud >= i {
		fmt.Print("San kitob ololmisan !")
		c.Mavjud -= i
	} else {
		fmt.Print("San kitob ololasan")
	}
}

func (c *book) Return(i int) {
	c.Mavjud += i
	fmt.Printf("%v book added %v\n", c.Nomi)
}

func (c *book) DisplayInfo() {
	fmt.Print("%+v\n", c)
}

func main() {

	newbook := book{

		Nomi:       "Sehrli uy",
		Aftori:     "Muhammad Yusuf",
		Mavjud:     5,
		Umumiysoni: 100,
	}
	newbook.borrow(3)
	newbook.DisplayInfo()
	newbook.Return(15)
}
