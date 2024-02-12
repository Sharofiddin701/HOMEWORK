package main

import "fmt"

func main() {

	var N int

	fmt.Print("Kun boshidan necha sekund o'tganini kiriting : ")
	fmt.Scan(&N)

	k := N / 60
	t := N % 60

	fmt.Println("Kun boshidan", k, "minut va ", t, "sekund vaqt o'tdi !!! \n")

}
