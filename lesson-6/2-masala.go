package main

import "fmt"

func harf(x string) string {

	var a string

	for i := len(x) - 1; i >= 0; i-- {

		a += string(x[i])

	}
	return a
}

func main() {

	var x string

	fmt.Print("Kiriting:")
	fmt.Scan(&x)

	fmt.Print(harf(x))

	harf(x)
}
