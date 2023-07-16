package main

import (
	"fmt"
)

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range ints {
		var status = ""

		if num%2 == 0 {
			status = "even"
		} else {
			status = "odd"
		}

		fmt.Println(num, "is", status)
	}

}
