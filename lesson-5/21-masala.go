package main

import "fmt"

func main() {

	var a int

	fmt.Print("a= ")
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Println("Yanvar oyida 29 ta kun bor")
	case 2:
		fmt.Println("Fevral oyida 30 ta kun bor")
	case 3:
		fmt.Println("Mart oyida 30 ta kun bor")
	case 4:
		fmt.Println("Aprel oyida 29 ta kun bor")
	case 5:
		fmt.Println("May oyida 30  ta kun bor")
	case 6:
		fmt.Println("Iyun oyida 30 ta kun bor")
	case 7:
		fmt.Println("Iyul oyida 27 ta kun bor")
	case 8:
		fmt.Println("Avgust oyida 28 ta kun bor")
	case 9:
		fmt.Println("Sentabr oyida 29 ta kun bor")
	case 10:
		fmt.Println("Oktabr oyida 30 ta kun bor")
	case 11:
		fmt.Println("Noyabr oyida 30 ta kun bor")
	case 12:
		fmt.Println("Dekabr oyida 31 ta kun bor")
	default:
		fmt.Println("Bunday fasl yo'q")

	}
}
