package main

import "fmt"

func main() {

	var N, target, nums int

	fmt.Print("Massiv uzunligini kiriting :")
	fmt.Scan(&N)

	fmt.Print("Target :")
	fmt.Scan(&target)

	for k := 0; k <= N; k++ {

		fmt.Print("a[", k, "]=")
		fmt.Scan(&nums)
	}
	for i := 0; i < N-1; i++ {

		for j = i + 1; j < N; j++ {

			if nums[i]+nums[j] == target {

				return i

			}
		}
	}

}
