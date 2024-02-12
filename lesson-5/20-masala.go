package main

import "fmt"

func main() {

	var a int

	fmt.Print("a= ")
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Println("Qish")
	case 2:
		fmt.Println("Qish")
	case 3:
		fmt.Println("Bahor")
	case 4:
		fmt.Println("Bahor")
	case 5:
		fmt.Println("Bahor")
	case 6:
		fmt.Println("Yoz")
	case 7:
		fmt.Println("Yoz")
	case 8:
		fmt.Println("Yoz")
	case 9:
		fmt.Println("Kuz")
	case 10:
		fmt.Println("Kuz")
	case 11:
		fmt.Println("Kuz")
	case 12:
		fmt.Println("Qish")
	default:
		fmt.Println("Bunday fasl yo'q")

	}
}
