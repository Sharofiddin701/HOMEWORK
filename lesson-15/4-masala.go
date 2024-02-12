package main

import "fmt"

func PrintValue(i interface{}) {

	if v, ok := i.(string); ok {
		fmt.Println("It is String", v)
	} else if v, ok := i.(int); ok {
		fmt.Println("It is Int", v)
	} else if v, ok := i.(float64); ok {
		fmt.Println("It is Float", v)
	} else if v, ok := i.(bool); ok {
		fmt.Println("It is Bool", v)
	} else {
		fmt.Println("Bunday yo'q")
	}
}
func main() {

	PrintValue(41)
	PrintValue("sdfsad")
	PrintValue(41.43)
	PrintValue(true)
}
