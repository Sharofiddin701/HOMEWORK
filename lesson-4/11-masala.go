package main

import "fmt"

func main() {

	var a, b int
	var k string

	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)

	fmt.Print("Sonlarga qaysi amalni qo'lamoqchisiz(*,/,+,-) ?:")
	fmt.Scan(&k)

	if k == "*" {
		fmt.Println("Ko'paytmasi=", a*b)
	} else if k == "+" {
		fmt.Println("Sonlarning yig'indisi=", a+b)
	} else if k == "-" {
		fmt.Println("Sonlarning ayirmasi=", a-b)
	} else {
		fmt.Println("Sonlarning bo'linmasi=", a/b)
	}
}
