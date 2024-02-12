package main

import "fmt"

func main() {

	var a float64
	pi := 3.14

	fmt.Println("a=")
	fmt.Scan(&a)

	k := 4 * pi * a * a
	l := 4 / 3 * pi * a * a * a

	fmt.Println("Sferaning yuzasi=", k)
	fmt.Println("Sferaning hajmi=", l)

}
