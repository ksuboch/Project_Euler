package main

import (
	"fmt"
)

func main() {
	fib1 := 1
	fib2 := 2
	fibCur := fib1 + fib2

	sum := 2

	for fibCur <= 4000000 {

		if fibCur%2 == 0 {
			sum += fibCur
		}

		fib1, fib2 = fib2, fibCur
		fibCur = fib1 + fib2
	}
	fmt.Println(sum)
}
