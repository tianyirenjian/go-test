package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	haystack := strings.Repeat("aaaa", 50000000) + "bb"
	needle := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabb"
	fmt.Println(len(haystack))
	t := time.Now()
	index := strStr(haystack, needle)
	fmt.Println(time.Since(t))
	fmt.Println(index)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if  needle == haystack {
		return 0
	}

	lenNeedle := len(needle)
	lenHaystack := len(haystack)

	if lenNeedle >= lenHaystack {
		return -1
	}

	needleHash := hash(needle)
	h := 0
	for i := 0; i <= lenHaystack -  lenNeedle; i ++ {
		if i == 0 {
			h = hash(haystack[i:i + lenNeedle])
		} else {
			h = newHash(h, haystack[i - 1], haystack[i + lenNeedle - 1])
		}
		if h == needleHash {
			if haystack[i:i + lenNeedle] == needle {
				return i
			}
		}
	}
	return -1
}

func hash(str string) int {
	hash := 0
	for _, c := range str {
		hash += int(c)
	}
	return hash
}

func newHash(hash int, minus uint8, plus uint8) int {
	return hash - int(minus) + int(plus)
}
