package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b>0 katta bo'lishi kerak !!! : ")
	fmt.Scan(&b)

	for i := 1; i < b; i++ {

		fmt.Println("i:", a)

	}

}
