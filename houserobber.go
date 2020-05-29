package main

import "fmt"

func main () {
	fmt.Println(rob([]int{1}))
	fmt.Println(rob([]int{1,2}))
	fmt.Println(rob([]int{1,2,4}))
	fmt.Println(rob([]int{2,7,9,3,1}))
	fmt.Println(rob([]int{902,563,377,287,120,555,900,644,667,601,487,621,589,947,245,280,394,963,970,498,944,482,262,419,379,469,195,663,557,667,939,251,184,958,624,460,253,813,448,852,943,936,799,647,245,537,475,395,907,810,978,712,391,884,434,893,609,747,693,366,104,871,963,563,617,188,447,185,675,282,388,798,925,538,750,192,957,183,477,102,415,596,346,747,841,968,611,415,240,821,951,194,975,819,571,193,229,752,730,130,226}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var dp []int
	dp = append(dp, nums[0], nums[1])
	if len(nums) >= 3 {
		dp = append(dp, nums[0] + nums[2])
		for i := 3; i < len(nums); i ++ {
			if nums[i] + dp[i - 2] > nums[i] + dp[i - 3] {
				dp = append(dp, nums[i] + dp[i - 2])
			} else {
				dp = append(dp, nums[i] + dp[i - 3])
			}
		}
	}
	if dp[len(dp) - 2] > dp[len(dp) - 1] {
		return dp[len(dp) - 2]
	}
	return dp[len(dp) - 1]
}
