package main

import "fmt"

func main() {

	var a int

	fmt.Print("a= ")
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Bunday hafta kuni yo'q")
	}
}
