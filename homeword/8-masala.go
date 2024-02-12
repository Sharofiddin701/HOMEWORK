package main

import (
	"fmt"
)

func main() {

	var a int

	// 1 va 2 soni kiritilmasin chunki 1 va 2 burchakli shakllar yo'q !

	fmt.Print("Shaklning tomonlari sonini kiriting:")
	fmt.Scan(&a)

	switch a / 10 {
	case 1:
		fmt.Print("o'n")
	case 2:
		fmt.Print("yigirma")
	case 3:
		fmt.Print("o'ttiz")
	case 4:
		fmt.Print("qirq")
	case 5:
		fmt.Print("ellik")
	case 6:
		fmt.Print("oltmish")
	case 7:
		fmt.Print("yetmish")
	case 8:
		fmt.Print("sakson")
	case 9:
		fmt.Print("to'qson")
	}

	switch a % 10 {
	case 0:
		fmt.Println("burchak")
	case 1:
		fmt.Println("birburchak")
	case 2:
		fmt.Println("ikkiburchak")
	case 3:
		fmt.Println("uchburchak")
	case 4:
		fmt.Println("to'rtburchak")
	case 5:
		fmt.Println("beshburchak")
	case 6:
		fmt.Println("oltiburchak")
	case 7:
		fmt.Println("yettiburchak")
	case 8:
		fmt.Println("sakkizburchak")
	case 9:
		fmt.Println("to'qqizburchak")
	}

}
