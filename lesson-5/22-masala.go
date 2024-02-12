package main

import "fmt"

func main() {

	var a, b float32
	var c int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)
	fmt.Print("Amalni kiriting:")
	fmt.Scan(&c)

	j := a + b
	h := a - b
	u := a / b
	p := a * b

	switch c {
	case 1:
		fmt.Println("Yig'imndisi :", j)
	case 2:
		fmt.Println("Ayirmasi :", h)
	case 3:
		fmt.Println("Bo'linmasi :", u)
	case 4:
		fmt.Println("Ko'paytmasi :", p)
	default:
		fmt.Println("Xato")
	}
}
