package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)

	fmt.Println(a, b)

	sonlar(&a, &b)

	fmt.Println(a, b)
}
func sonlar(a, b *int) {

	*a, *b = *b, *a

}
