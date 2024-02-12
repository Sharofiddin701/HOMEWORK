package main

import "fmt"

func main() {

	var W int

	fmt.Print("a= ")
	fmt.Scan(&W)

	if W > 0 {
		fmt.Println("Son musbat ", W+1)
	} else {
		fmt.Println("Son manfiy: ", W+2)
	}
}
