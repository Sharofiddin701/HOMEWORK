package main

import "fmt"

func main() {

	var harf string

	fmt.Print("Harfni kitiriting : ")
	fmt.Scan(&harf)

	k := (("a" <= harf) && (harf <= "z"))

	fmt.Printf("Unli: %v\n", k)
}
