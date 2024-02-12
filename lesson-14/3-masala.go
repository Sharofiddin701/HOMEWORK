package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	k := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	v := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(k); i++ {
		if (coordinates[0] == k[i][0]) && (int(coordinates[1]-'0') == v[i]) {
			return (i%2 == 0 && v[i]%2 == 1) || (i%2 == 1 && v[i]%2 == 0)
		}
	}
	return false
}

func main() {
	var coordinates string

	fmt.Println("Kiriting:")
	fmt.Scan(&coordinates)

	if squareIsWhite(coordinates) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
