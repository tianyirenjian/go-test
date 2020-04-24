package main

import (
	"fmt"
	"sort"
	"strings"
)

func main () {
	words := []string{"leetcode","et" ,"code", "let"}
	fmt.Println(stringMatching(words))
}

type stringArray []string

func (s stringArray) Len() int {
	return len(s)
}

func (s stringArray) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s stringArray) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func stringMatching(words []string) []string {
	sort.Sort(stringArray(words))
	length := len(words)
	words2 := []string{}
	for i := 0; i < length; i ++ {
		for j := i + 1; j < length; j ++ {
			if strings.Contains(words[j], words[i]) {
				words2 = append(words2, words[i])
				break
			}
		}
	}
	return words2
}
