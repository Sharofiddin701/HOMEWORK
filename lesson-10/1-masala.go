package main

import (
	"fmt"
	"strings"
)

func main() {

	sentense := " I have a macbook"
	subsstring := "macbook"

	a := strings.Index(sentense, subsstring)

	fmt.Println(a)

}
