package main

import "fmt"

func FindOutlier(integers []int) int {
	var oddCount = 0
	var evenCount = 0

	var lastOdd = -1
	var lastEven = -1

	for _, val := range integers {
		if val%2 == 0 {
			evenCount++
			lastEven = val
		} else {
			oddCount++
			lastOdd = val
		}

		if evenCount > 1 && lastOdd != -1 {
			return lastOdd
		} else if oddCount > 1 && lastEven != -1 {
			return lastEven
		}
	}

	return 0
}

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
}
