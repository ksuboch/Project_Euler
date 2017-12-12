package main

import (
	"fmt"
	"math"
)

const val = 600851475143

func primes(n int64) []int64 {
	i := int64(2)
	cur := n
	maxI := delimeterLimit(n)
	res := []int64{}
	for i <= maxI {
		if cur%i == 0 && isPrime(i) {
			res = append(res, i)
			cur = n / i
			maxI = delimeterLimit(cur)
		} else {
			i++
		}
	}
	return res
}

func isPrime(n int64) bool {
	i := int64(2)
	maxI := delimeterLimit(n)
	for i <= maxI {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

func delimeterLimit(n int64) int64 {
	return int64(math.Sqrt(float64(n)))
}

func main() {
	fmt.Println(primes(val))
}
