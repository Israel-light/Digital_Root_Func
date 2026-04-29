package main

import (
	"fmt"
)

func singleDigit(n int) int {
	if n < 0 {
		n = -n
	}
	if n < 10 {
		return n
	}

	num := fmt.Sprintf("%d", n)

	addition := 0
	if n > 9 {
		for i := 0; i < len(num); i++ {
			addition += int(num[i]) - '0'
		}
	}

	if addition > 9 {
		return singleDigit(addition)
	}

	return addition
}

func main() {
	fmt.Println(singleDigit(467))
	fmt.Println(singleDigit(21))
	fmt.Println(singleDigit(1000))
	fmt.Println(singleDigit(1))
	fmt.Println(singleDigit(-36))
}
