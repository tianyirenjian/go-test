package main

import "fmt"

func main() {
	nums :=  [][]int{
		{1,2,2,6,3,6,1,8,9,4,7,6,5,6,8,2,6,1,3,6,6,6,3,2,4,9,4,5,9,8,2,2,1,6,1,6,2,2,6,1,8,6,8,3,2,8,5,8,0,1,4,8,7,9,0,3,9,4,8,0,2,2,5,5,8,6,3,1,0,2,4,9,8,4,4,2,3,2,2,5,5,9,3,2,8,5,8,9,1,6,2,5,9,9,3,9,7,6,0,7,8,7,8,8,3,5,0},
		{2,3,1,1,4},
		{3,2,1,0,4},
		{1},
		{2, 0, 0},
	}
	for _, n := range nums {
		fmt.Println(canJump(n))
	}
}

func canJump(nums []int) bool {
	if len(nums) == 1 || !hasZero(nums) {
		return true
	}
	length := len(nums)
	for i := 0; i < length - 1; i ++ {
		if nums[i] == 0 {
			has := false
			for j := i - 1; j >= 0; j -- {
				if nums[j] + j > i {
					has = true
					break
				}
			}
			if !has {
				return false
			}
		}
	}
	return true
}

func hasZero(nums []int) bool {
	length := len(nums)
	for i := 0; i < length; i ++ {
		if nums[i] == 0 {
			return true
		}
	}
	return false
}
