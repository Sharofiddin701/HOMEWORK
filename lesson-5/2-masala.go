package main

import "fmt"

func main() {

	var W int

	fmt.Print("a= ")
	fmt.Scan(&W)

	if ((W%4 == 0) && (W%100 != 0)) || (W%400 == 0) {
		fmt.Println("Kabisa yili")
	} else {
		fmt.Println("Kabisa yili emas !!!")
	}
}
