package main

import "fmt"

func main () {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(9))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(122252221))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x >=0 && x <=9 {
		return true
	}
	z := x
	y := 0
	for x > 0 {
		y = y * 10 + x % 10
		x = x / 10
	}
	return y == z
}

