package main

import "fmt"

func main() {

	var target int
	nums := []int{2, 7, 11, 15}

	fmt.Println(nums)

	fmt.Print("target:")
	fmt.Scan(&target)

	for i := 0; i <= 3; i++ {

		for j := i + 1; j <= 3; j++ {

			if nums[i]+nums[j] == target {

				fmt.Print(i, j,"\n")
			}

		}

	}

}
