package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Student struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Course   int    `json:"course"`
	Contract int    `json:"contract"`
}
func main() {
	file, err := os.Open("fayl.json")
	if err != nil {
		fmt.Println("Open erroring file", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var (
		fayls []Student
	)
	err = decoder.Decode(&fayls)

	if err != nil {
		fmt.Println("Error decoding json single obj:", err)
		return
	}
	fmt.Printf("Decoded fayl :%+v", fayls)

	for _, student := range fayls {
		a := student.Name

		if checkFileExists(a) {

			b := a + strconv.Itoa(rand.Intn(100))

			file, err := os.Create(b)
			if err != nil {
				fmt.Println("Error opening file", err)
				return
			}
			defer file.Close()
			fmt.Println("File open successfully")
		} else {

			file, err := os.Create(a)
			if err != nil {
				fmt.Println("Error opening file", err)
				return
			}
			defer file.Close()
			fmt.Println("File open successfully")
		}
	}

}
func checkFileExists(filepath string) bool {
	_, error := os.Stat(filepath)
	return !errors.Is(error, os.ErrNotExist)

}
