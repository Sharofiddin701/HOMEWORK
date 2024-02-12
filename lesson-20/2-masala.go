package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var son int
	var name, gmail, job string
	file, err := os.Create("namejob.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Print("Enter the number of entries: ")
	fmt.Scan(&son)

	for i := 0; i < son; i++ {
		fmt.Printf("Entry %d:\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Email: ")
		fmt.Scan(&gmail)

		if strings.HasSuffix(gmail, "@gmail.com") || strings.HasSuffix(gmail, "@mail.ru") {
		continue
		} else {
			fmt.Println("Email must end with '@gmail.com' or '@mail.ru'.")
			fmt.Print("Job: ")
			fmt.Scan(&job)
		}

		line := fmt.Sprintf("Name: %s, Email: %s, Job: %s\n", name, gmail, job)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Data written successfully to namejob.txt")
}
