package main

import (
	"fmt"
	"time"
)

func main() {
	var a, b int

	fmt.Print("Sonni kiriting :")
	fmt.Scan(&a)
	fmt.Print("Sekendni kiritign:")
	fmt.Scan(&b)

	if b > 60 {
		fmt.Println("60 dan kichikroq son kiriting !")
		return

	}
	for i := 1; i <= a; i++ {

		fmt.Println(i)
		time.Sleep(time.Duration(time.Second * time.Duration(b)))
	}

}
