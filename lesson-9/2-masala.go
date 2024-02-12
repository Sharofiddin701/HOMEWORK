package main

import "fmt"

func main() {
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 1, 3}))
}

func birthdayCakeCandles(candles []int32) int32 {
	c := candles[0]
	var counter int32
	for _, i := range candles {
		if c < int32(i) {
			c = int32(i)
		}
	}

	for v := 0; v <= len(candles)-1; v++ {
		if c == candles[v] {
			counter++
		}
	}
	return int32(counter)
}
