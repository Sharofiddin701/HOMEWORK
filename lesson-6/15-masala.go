package main

import "fmt"

func son(a int) {

	for i := 1; i <= a; i++ {

		if a%i == 0 {

			fmt.Println(i)

		}
	}

}
func main() {

	var a int

	fmt.Print("A=")
	fmt.Scan(&a)

	son(a)

}
