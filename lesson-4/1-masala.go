package main

import "fmt"

func main() {

	k := ((true && false) || !(true || false)) && ((true && false) && !(true || false))

	fmt.Printf("Unli: %v\n", k)
}
