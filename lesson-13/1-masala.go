package main

import "fmt"

func main() {

	var a int = 42

	var ptr *int

	ptr = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println(&ptr)

}
