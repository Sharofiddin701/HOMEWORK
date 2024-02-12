package main

import "fmt"

func soz(a, b int, t string) string {

	switch t {
	case "qo'shish":
		fmt.Println("Yig'indining qiymati:", a+b)
	case "Ayirish":
		fmt.Println("Ayirmaning qiymati:", a-b)
	case "Ko'paytirish":
		fmt.Println("Ko'paytmaning qiymati:", a*b)
	case "Bo'lish":
		fmt.Println("Bo'linmaning qiymati:", a/b)
	}
	return t
}
func main() {
	var a, b int
	var t string

	fmt.Print("A:")
	fmt.Scan(&a)
	fmt.Print("B:")
	fmt.Scan(&b)
	fmt.Print("Amalni kiriting:")
	fmt.Scan(&t)

	fmt.Println(soz(a, b, t))

}
