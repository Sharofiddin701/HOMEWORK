package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {

	harf := strings.ToLower(s)

	return harf
}

func main() {

	var s string

	fmt.Print("s:")
	fmt.Scan(&s)

	fmt.Println(toLowerCase(s))
}
