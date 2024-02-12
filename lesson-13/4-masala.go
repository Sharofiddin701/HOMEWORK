package main

import (
	"fmt"
	"strconv"
)

func main() {

	text := "qwert"

	number, err := strconv.Atoi(text)
	if err != nil {
		panic("Invailid number")
	}

	fmt.Println("sadsas", number)

}
