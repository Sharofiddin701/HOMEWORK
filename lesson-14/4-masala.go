package main

import "fmt"

func interpret(command string) string {
	str := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			str += "G"
		} else if command[i] == '(' && command[i+1] == ')' {
			str += "o"
		} else if command[i] == '(' && command[i+1] == 'a' {
			str += "al"
		}
	}
	return str
}

func main() {
	var command string
	fmt.Scan(&command)

	res := interpret(command)
	fmt.Println(res)

}
