package main

import "fmt"

func main() {

	k := ((111%11 == 1 && 1*0 == 0) || !(0/1 == 1) || (3695/2 == 1847))

	fmt.Printf("Unli: %v\n", k)
}
