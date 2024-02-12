package main

import "fmt"

func main() {

	var a int

	fmt.Print("a= ")
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Println("Yomon")
	case 2:
		fmt.Println("Qoniqarsiz")
	case 3:
		fmt.Println("Qoniqarli")
	case 4:
		fmt.Println("Yaxshi")
	case 5:
		fmt.Println("A'lo")
	default:
		fmt.Println("Xato")
	}
}
