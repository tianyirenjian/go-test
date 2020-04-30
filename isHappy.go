package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	for i := 0; i < 10000000; i ++ {
		isHappy(20)
	}
	fmt.Println(time.Since(t))
}

func isHappy(n int) bool {
	cache := make(map[int]int, 0)
	for n != 1 {
		a := 0
		for n > 0 {
			a += (n % 10) * (n % 10)
			n = n / 10
		}
		if a == 1 {
			return true
		}

		if _, ok := cache[a]; ok {
			return false
		}
		cache[a] = 1
		n = a
	}
	return true
}
