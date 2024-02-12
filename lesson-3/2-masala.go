package main

import "fmt"

func main() {

	var name string
	var age int
	var work string

	fmt.Print("Mening ismim :")
	fmt.Scan(&name)
	fmt.Print("Mening yoshim :")
	_, err := fmt.Scan(&age)
	fmt.Print("Mening ishim :")
	fmt.Scan(&work)
	fmt.Println(err)

	k:=fmt.Sprintf("Mening ismim : %v yoshim %v  ishim %v\n ", name, age, work)

	fmt.Print("Sprintf",k)
}
	//fmt.Printf("Mening ismim : %v yoshim %v  ishim %v ", name, age, work)
// }
