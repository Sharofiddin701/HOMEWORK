package main

import "fmt"

func main() {

	var a, b, c, d, f, k int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)
	fmt.Print("c= ")
	fmt.Scan(&c)

	if a > 0 {
		a = 1
	} else {
		d = 1
	}

	if b > 0 {
		b = 1
	} else {
		f = 1
	}

	if c > 0 {
		c = 1
	} else {
		k = 1
	}
	fmt.Println("ta musbat son bor", a+b+c)
	fmt.Println("ta manfiy son bor", d+f+k)

}
