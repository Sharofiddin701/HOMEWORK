package main

import "fmt"
		"math"

func main() {

	var W, H float32

	fmt.Print("Odamning og'irligini kiriting : ")
	fmt.Scan(&W)
	fmt.Print("Odamning bo'y uzunligini kiriting : ")
	fmt.Scan(&H)

	k := W / (H * H)

	fmt.Println(" K=", k)

	r := (k < 18.5)
	y := ((18.5 < k) && (k < 25.0))
	o := ((25.0 < k) && (k < 30.0))
	p := (k > 30.0)

	fmt.Println("Kam vaznli ", r)
	fmt.Println("Sog'lom ", y)
	fmt.Println("Ortiqcha vaznga ega ! ", o)
	fmt.Println("Semirib ketgan ", p)
}
