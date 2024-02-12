package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("File.txt")
	if err != nil {
		fmt.Println("Xatolik: Faylni ochib bo'lmadi:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	var words []string
	for scanner.Scan() {
		line := scanner.Text()

		for _, word := range strings.Fields(line) {
			if num, err := strconv.Atoi(word); err == nil {
				nums = append(nums, num)
			} else {
				words = append(words, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik: Faylni o'qib bo'lmadi:", err)
		return
	}

	fmt.Println("Sonlar:", nums)
	fmt.Println("So'zlar:", words)
}
