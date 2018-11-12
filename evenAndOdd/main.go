package main

import (
	"fmt"
)

func main() {
	sliceOfInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range sliceOfInts {
		if sliceOfInts[i]%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
