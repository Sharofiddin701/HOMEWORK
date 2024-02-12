package main

import "fmt"

func son(a int) bool {

	k := 0
	tu := true

	for i := 1; i <= a; i++ {

		k += i

		if k%2 == 0 {
			tu = true
		} else {
			tu = false
		}
	}
	return tu
}
func main() {

	var a int

	fmt.Print("N=")
	fmt.Scan(&a)

	fmt.Println(son(a))

}
