package main

import (
	"errors"
	"fmt"
)

func twonumber(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("Nolga teng bo'lishi mumkin emas !!!")
	}
		return x / y, nil

}

func main() {

	var f, g int
	fmt.Print("x=")
	fmt.Scan(&f)
	fmt.Print("y=")
	fmt.Scan(&g)

	twonumber(f, g)
	fmt.Print(twonumber)

}
