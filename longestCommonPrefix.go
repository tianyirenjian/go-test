package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"good", "god", "good for"}))
}

func longestCommonPrefix(strs []string) string {
	count := 0
	length := len(strs)
	if length == 0 {
		return ""
	}
	for true {
		b := true
		var c uint8
		for i := 0; i < length; i ++ {
			if i == 0 {
				if len(strs[i]) <= count {
					b = false
					break
				}
				c = strs[i][count]
			} else {
				if len(strs[i]) <= count || strs[i][count] != c {
					b = false
					break
				}
			}
		}
		if !b {
			break
		} else {
			count ++
		}
	}
	return strs[0][0:count]
}
