package main

import "fmt"

func main() {

	var words string

	fmt.Print("Kiriting:")
	fmt.Scan(&words)

	fmt.Print(firstPalindrome)

}

func firstPalindrome(words []string) string {

	for i := 0; i <= len(words); i++ {

		for j := len(words); j >= 0; j-- {

			if words[i] == words[j] {

				fmt.Print(words)

			}

		}

	}
}
