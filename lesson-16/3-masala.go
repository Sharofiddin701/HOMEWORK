package main

import (
	"errors"
	"fmt"
)

func twonumber(x, y int) (int, error) {

	if x > 0 && y > 0 {

		return x / y, nil

	} else if y == 0 {

		return 0, errors.New("Nolga teng bo'lishi mumkin emas !!!")
	}

}

func main() {

	var f, g int
	fmt.Print("x=")
	fmt.Scan(&f)
	fmt.Print("y=")
	fmt.Scan(&g)

	twonumber(f, g)

}
