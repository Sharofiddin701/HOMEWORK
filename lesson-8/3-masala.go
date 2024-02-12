package main

import "fmt"

func main() {

	sonlar := []int{1, 2, 3, 4, 5, 1, 6, 3, 1}

	newmap := make(map[int]bool)

	newarr := make([]int, 0, 10)

	for _, number := range sonlar {

		if !newmap[sonlar] {

			sonlar = append(newarr, number)
			newmap[number] = true

		}
		fmt.Println(newarr)
	}

}
